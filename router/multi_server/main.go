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

	// Define a basic route handling both GET and POST requests
	r.HandleFunc("/", "GET POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		// Detect which server (host/port) received the request
		serverHost := req.Host

		// Generate a dynamic response showing the active server
		response := fmt.Sprintf(`
			<h1>Hello World</h1>
			<p>Successfully connected!</p>
			<p>Served by listener: <strong style="color: green;">%s</strong></p>
		`, serverHost)

		w.Write([]byte(response))
	})

	// Define multiple listeners (e.g., different ports or interfaces)
	listeners := router.Listeners{
		{Addr: "localhost:8000"},
		{Addr: "localhost:8001"},
	}

	// Start the server on all defined listeners simultaneously

	if err := r.MultiListenAndServe(listeners); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
