package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// --- BASIC CUSTOM 404 HANDLER ---
	// The NotFound method allows you to override the default "404 page not found" message.
	// This is useful for returning simple custom text, localized messages, or
	// lightweight inline HTML without managing external files.
	r.NotFound(func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		// Set the header to indicate the content type
		w.Header().Set("Content-Type", "text/html")

		// It is important to explicitly set the 404 status code
		w.WriteHeader(http.StatusNotFound)

		// Send your custom message
		w.Write([]byte(`<h1>404</h1><p>The page you are looking for does not exist.</p>`))
	})

	// Standard route to test the navigation
	r.GET("/", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`Hello World! <br> <a href="/any-broken-link">Click here to trigger the 404 handler</a>`))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
