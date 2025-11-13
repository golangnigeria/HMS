package render

import (
	"bytes" // ✅ Added for toJson
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData injects default values into template data
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders HTML templates using Go's html/template package
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser:", err)
	}
}

// CreateTemplateCache parses all page, layout, and partial templates
// and stores them in a cache (map) for efficient reuse.
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(pathToTemplates, "/pages/**/", "*.page.html"))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// ✅ Attach custom functions here
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Parse layouts
		matches, err := filepath.Glob(filepath.Join(pathToTemplates, "/layouts/", "*.layout.html"))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseFiles(matches...)
			if err != nil {
				return myCache, err
			}
		}

		// Parse partials
		partials, err := filepath.Glob(filepath.Join(pathToTemplates, "/partials/", "*.partial.html"))
		if err != nil {
			return myCache, err
		}
		if len(partials) > 0 {
			ts, err = ts.ParseFiles(partials...)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
