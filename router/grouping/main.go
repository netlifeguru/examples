package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	// Initialize a new router instance
	r := router.New()
	// Create a new route group with the prefix "/api/v1"
	// Route groups are perfect for versioning your API or applying
	// specific middlewares (like Authentication) only to a subset of routes.
	api := r.Group("/api/v1")
	// Using a block { ... } is purely for visual organization and scoping,
	// but it is highly recommended to keep your routing logic clean.
	{
		// Resulting URL: GET or POST /api/v1/
		api.HandleFunc("/", "GET POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusCreated) // 201 Created
			w.Write([]byte(`<h1>/api/v1/</h1>`))
		})

		// Resulting URL: GET or POST /api/v1/user/{id}
		// The group prefix is automatically prepended to the route path.
		api.HandleFunc("/user/{id}", "GET POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
			userID := ctx.Param("id")
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusCreated) // 201 Created
			w.Write([]byte(fmt.Sprintf(`<h1>/api/v1/user:%s</h1>`, userID)))
		})
	}

	// Start the server on port 8000
	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
