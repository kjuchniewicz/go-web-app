package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf dodaje ochronę CSRF do wszystkich próśb POST
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad ładuje i zapisuje sesje dla kazdej prośby
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
