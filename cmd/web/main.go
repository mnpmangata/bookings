package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mnpmangata/bookings/pkg/config"
	"github.com/mnpmangata/bookings/pkg/handlers"
	"github.com/mnpmangata/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.Appconfig
var session *scs.SessionManager

func main() {

	//Change this to true when in production, false for Localhost:8080
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //should the cookies persist after the browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	//When true: use template cache, when false: rebuild cache and read from disk (use on devlmnt mode)
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
