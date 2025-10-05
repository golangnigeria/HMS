package handlers

import (
	"net/http"

	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
)

func (m *Repository) DoctorsLoginPage(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "doctorLogin.page.html", td)
}

func (m *Repository) DoctorsRegistrationPage(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "doctorRegistration.page.html", td)
}
