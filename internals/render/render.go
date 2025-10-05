package render

import (
	"bytes"
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
var functions = template.FuncMap{}

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
		// Use the cached templates
		tc = app.TemplateCache
	} else {
		// Rebuild cache for every request (dev mode)
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
	// Initialize a map to store compiled templates.
	// The key will be the filename (e.g., "home.page.html"), 
	// and the value will be the parsed *template.Template.
	myCache := map[string]*template.Template{}

	// Find all files that match "*.page.html" inside the "pages" folder
	// This is where our individual page templates (home, about, contact, etc.) live.
	pages, err := filepath.Glob(filepath.Join(pathToTemplates, "/pages/", "*.page.html"))
	if err != nil {
		// If there’s an error while searching for page templates, return immediately
		return myCache, err
	}

	// Loop through each page template found
	for _, page := range pages {
		// Extract only the filename (e.g., "home.page.html") 
		// to use as the map key later.
		name := filepath.Base(page)

		// Start a new template with the given filename,
		// attach any custom functions via `.Funcs(functions)`,
		// and parse the current page file.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			// If parsing the page fails, stop and return the error
			return myCache, err
		}

		// --- Parse layouts ---
		// Layout templates are stored in the "layouts" folder (e.g., base layout).
		// They typically include shared HTML (header, footer, etc.)
		matches, err := filepath.Glob(filepath.Join(pathToTemplates, "/layouts/", "*.layout.html"))
		if err != nil {
			return myCache, err
		}

		// If layouts exist, parse and attach them to the current template
		if len(matches) > 0 {
			ts, err = ts.ParseFiles(matches...)
			if err != nil {
				return myCache, err
			}
		}

		// --- Parse partials ---
		// Partials are reusable template snippets (e.g., navbar, sidebar, footer).
		// They live inside the "partials" folder.
		partials, err := filepath.Glob(filepath.Join(pathToTemplates, "/partials/", "*.partial.html"))
		if err != nil {
			return myCache, err
		}

		// If partials exist, parse and attach them as well
		if len(partials) > 0 {
			ts, err = ts.ParseFiles(partials...)
			if err != nil {
				return myCache, err
			}
		}

		// Add the fully parsed template (page + layout + partials) to the cache map.
		myCache[name] = ts
	}

	// Return the completed template cache map
	return myCache, nil
}

