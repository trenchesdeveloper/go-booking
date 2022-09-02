package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/trenchesdeveloper/go-bookings/pkg/config"
	"github.com/trenchesdeveloper/go-bookings/pkg/handlers"
	"github.com/trenchesdeveloper/go-bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false
	// create session
	session = scs.New()

	// set session lifetime
	session.Lifetime = 24 * time.Hour

	// set session cookie name
	session.Cookie.Name = "session"

	// persist session data across multiple requests
	session.Cookie.Persist = true

	// set session same site policy
	session.Cookie.SameSite = http.SameSiteLaxMode

	// set session secure flag
	session.Cookie.Secure = app.InProduction

	// set session to app config
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// set the app config in the render package
	render.NewTemplates(&app)

	fmt.Println("Server is running on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Fatal(srv.ListenAndServe())
}
