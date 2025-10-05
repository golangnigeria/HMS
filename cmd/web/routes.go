package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/handlers"

	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	/* Doctors routes */
	mux.Get("/login/role/doctor", handlers.Repo.DoctorsLoginPage)
	mux.Get("/registration/role/doctor", handlers.Repo.DoctorsRegistrationPage)

	/* Patients routes */
	mux.Get("/registration/role/patient", handlers.Repo.PatientsRegistrationPage)
	mux.Get("/login/role/patient", handlers.Repo.PatientsLoginPage)

	/* Lab/clinic routes */
	mux.Get("/login/role/lab", handlers.Repo.LabLoginPage)
	mux.Get("/registration/role/lab", handlers.Repo.LabRegistrationPage)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
