package handlers

import (
	"net/http"

	"github.com/mnpmangata/bookings/pkg/config"
	"github.com/mnpmangata/bookings/pkg/models"
	"github.com/mnpmangata/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.Appconfig
}

// NewRepo creates a new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home takes a function receiver to access the repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About takes a function receiver to access the repository
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
