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
	mux.Group(func(r chi.Router) {
		r.Get("/", handlers.Repo.Home)
		r.Get("/about", handlers.Repo.About)
	})

	// --- Doctor Routes ---
	mux.Route("/doctor", func(r chi.Router) {
		r.Get("/login", handlers.Repo.DoctorsLoginPage)
		r.Get("/register", handlers.Repo.DoctorsRegistrationPage)
	})

	// --- Patient Routes ---
	mux.Route("/patient", func(r chi.Router) {
		r.Get("/login", handlers.Repo.PatientsLoginPage)
		r.Get("/register", handlers.Repo.PatientsRegistrationPage)
	})

	// --- Laboratory Routes ---
	mux.Route("/lab", func(r chi.Router) {
		r.Get("/login", handlers.Repo.LabLoginPage)
		r.Get("/register", handlers.Repo.LabRegistrationPage)
		r.Get("/portal", handlers.Repo.LabPortalPage)
		r.Get("/tests", handlers.Repo.LabPortalPage) // optional alias
	})

	// --- Static Files ---
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
