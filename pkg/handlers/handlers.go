package handlers

import (
	"fmt"
	"net/http"

	"github.com/krshnas/myapp/pkg/config"
	"github.com/krshnas/myapp/pkg/models"
	"github.com/krshnas/myapp/pkg/render"
)

// Repo the repository used by the handlers
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
func (h *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	h.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (h *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := h.App.Session.Get(r.Context(), "remote_ip")
	fmt.Println(remoteIp)
	if remoteIpStr, ok := remoteIp.(string); ok {
		stringMap["remote_ip"] = remoteIpStr
	} else {
		stringMap["remote_ip"] = ""
	}

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
