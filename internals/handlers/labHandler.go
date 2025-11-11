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

func (m *Repository) LabPortalPage(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Name   string
		Role   string
		Avatar string
	}{
		Name:   "Dr. Ada Lovelace",
		Role:   "Lab Technician",
		Avatar: "/static/images/default-avatar.png",
	}

	tests := []models.Test{
		{Name: "Complete Blood Count", Category: "Hematology", Price: "₦5,000", Status: "Active"},
		{Name: "Urinalysis", Category: "Chemistry", Price: "₦3,000", Status: "Inactive"},
		{Name: "Liver Function Test", Category: "Biochemistry", Price: "₦8,500", Status: "Active"},
	}

	

	pagination := struct {
		Total int
		Start int
		End   int
		Prev  string
		Next  string
	}{
		Total: 100,
		Start: 1,
		End:   10,
		Prev:  "prev",
		Next:  "next",
	}

	dataMap := map[string]interface{}{
		"tests": tests,
		"pagination": pagination,
	}

	data := &models.TemplateData{
		User:        user,
		Active: "lab_portal", 
		Data:        dataMap,
	}

	render.RenderTemplate(w, r, "lab_portal.page.html", data)
}
