package handlers

import (
	"net/http"

	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
)

func (m *Repository) AdminBlogCreateForm(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	render.RenderTemplate(w, r, "blog_create.page.html", td)
}
