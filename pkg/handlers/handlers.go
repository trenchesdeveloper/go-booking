package handlers

import (
	"github.com/trenchesdeveloper/go-bookings/pkg/config"
	"github.com/trenchesdeveloper/go-bookings/pkg/models"
	"github.com/trenchesdeveloper/go-bookings/pkg/render"
	"net/http"
)

// Repo is the repository type
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["key1"] = "value1"
	stringMap["key2"] = "value2"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
