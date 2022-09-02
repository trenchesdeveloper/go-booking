package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		next.ServeHTTP(w, r)
	})
}

// NoSurf is a middleware that blocks requests with unsafe methods csrf
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
	})

	return csrfHandler
}

// SessionLoader is a middleware that loads the session from the cookie
func SessionLoader(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
