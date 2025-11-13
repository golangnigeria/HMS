package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/handlers"
)

// SetupRouter configures and returns the main router
func SetupRouter(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// --- Global Middleware ---
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// --- Public Routes ---
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	 
	// --- Blog Routes ---
	mux.Group(func(r chi.Router) {
		r.Get("/blog/bloglist", handlers.Repo.BlogList)
		r.Get("/blog/{slug}", handlers.Repo.BlogSingle)
	})

	// --- Admin Routes ---
	mux.Group(func(r chi.Router) {
		r.Get("/admin/blog/create", handlers.Repo.AdminBlogCreateForm)
	})


	// --- Static Files ---
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
