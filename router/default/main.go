package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	// Initialize a new router instance
	r := router.New()

	// Define a basic route handling both GET and POST requests
	r.HandleFunc("/", "GET POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusCreated) // 201 Created
		w.Write([]byte(`<h1>Hello World</h1>`))
	})

	r.HandleFunc("/user/{id}", "GET", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		//ctx.Param("id")
	})

	// Start the server on port 8000
	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
