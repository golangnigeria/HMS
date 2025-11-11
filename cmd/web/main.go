package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/driver"
	"github.com/golangnigeria/kinicart/internals/handlers"
	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
	"github.com/joho/godotenv"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application entry
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// What i am going to put in the session
	gob.Register(models.User{})

	// change true when in production
	inProduction, err := strconv.ParseBool(os.Getenv("IN_PRODUCTION"))
	if err != nil {
		inProduction = false // Default to false if there's an error
	}
	app.InProduction = inProduction

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.UseCache = app.InProduction

	// connect to database
	log.Println("connecting to database...")
	dsn := os.Getenv("DATABASE_URL")
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	defer db.SQL.Close()

	log.Println("Connected to database!")


	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// storing it in the cache
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = ":8080" 
	}
	fmt.Printf("Starting application on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: SetupRouter(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
