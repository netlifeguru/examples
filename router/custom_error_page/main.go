package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// Serving static assets (CSS, JS, Images)
	// This ensures your 404 page can load 'assets/style.css' correctly.
	r.Static("/assets/", "./public")

	// --- CUSTOM 404 NOT FOUND HANDLER ---
	// Instead of hardcoding HTML strings, we load a dedicated file.
	// This makes it easier to manage themes and marketing content.
	r.NotFound(func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")

		// Try to read the 404 template from your local directory
		htmlContent, err := os.ReadFile("./templates/404.html")
		if err != nil {
			// Fallback if the file is missing
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Page Not Found (and template missing)"))
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(htmlContent)
	})

	// Basic route to test the application
	r.GET("/", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<h1>Welcome to netlife.guru</h1><p>Try visiting a non-existent URL to see the custom 404.</p>`))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
