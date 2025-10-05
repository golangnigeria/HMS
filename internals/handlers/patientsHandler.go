package handlers

import (
	"net/http"

	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
)

func (m *Repository) PatientsRegistrationPage(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "patientRegistration.page.html", td)
}


func (m *Repository) PatientsLoginPage(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "patientLogin.page.html", td)
}