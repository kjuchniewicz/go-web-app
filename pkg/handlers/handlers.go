package handlers

import (
	"net/http"

	"github.com/kjuchniewicz/go-web/models"
	"github.com/kjuchniewicz/go-web/pkg/config"
	"github.com/kjuchniewicz/go-web/pkg/render"
)

// Repo repozytorium używane przez programy obsługi
var Repo *Repository

// Repository to typ repozytorium
type Repository struct {
	App *config.AppConfig
}

// NweRepo tworzy nowe repozytorium
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers ustawia repozytorium dla programów obsługi
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Siemka, znowu."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
