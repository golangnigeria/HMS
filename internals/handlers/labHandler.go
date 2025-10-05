package handlers

import (
	"net/http"

	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
)

func (m *Repository) LabRegistrationPage(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "labRegistration.page.html", td)
}


func (m *Repository) LabLoginPage(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "labLogin.page.html", td)
}