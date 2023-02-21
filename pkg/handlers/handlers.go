package handlers

import (
	"net/http"

	"github.com/ZhijiunY/golang-web/pkg/config"
	"github.com/ZhijiunY/golang-web/pkg/models"
	"github.com/ZhijiunY/golang-web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// call renderTemplate
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again, and again and again!"

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
