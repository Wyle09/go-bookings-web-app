package handler

import (
	"net/http"

	"github.com/wyle09/go-bookings-web-app/pkg/config"
	"github.com/wyle09/go-bookings-web-app/pkg/models"
	"github.com/wyle09/go-bookings-web-app/pkg/render"
)

// Repository used by the handlers
var Repo *Repository

// Is the repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (receiver *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "home.page.gohtml", &models.TemplateData{})
}

// About page handler
func (receiver *Repository) About(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{})
}
