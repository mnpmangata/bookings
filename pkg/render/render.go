package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/mnpmangata/bookings/pkg/config"
	"github.com/mnpmangata/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.Appconfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.Appconfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	//app.UseCache is true
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		//rebuild cache and read from the disk
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer) //create a buffer

	td = AddDefaultData(td)

	_ = t.Execute(buf, td) //execute t, pass value of td to buf

	_, err := buf.WriteTo(w) //ignore number of bytes
	if err != nil {
		fmt.Println("Error writing template to browser:", err)
	}
}

// Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	//load all templates to cache
	myCache := map[string]*template.Template{}

	//go to templates folder and find anything that ends with .page.tmpl
	//Glob returns names of files with the file pattern
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) //get the name of the page only from the filepath
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//find a file with .layout.tmpl on folder templates
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		//parse layout if matches > 0
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
} //end of CreateTemplateCache()
