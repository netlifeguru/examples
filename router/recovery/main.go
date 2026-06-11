package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// Custom Recovery Handler
	// This function is triggered whenever a panic occurs in any of your route handlers.
	// It allows you to return a clean, custom error page or JSON response to the user.
	//
	// FAIL-SAFE FEATURE: The router has an internal secondary recovery mechanism!
	// If this custom Recovery function itself panics (e.g., nil pointer dereference here),
	// the router's internal fail-safe will catch it, prevent the server from crashing,
	// and automatically return a standard 500 Internal Server Error.
	// The application will never hang.
	r.Recovery(func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {

		// You can safely handle the error here.
		// Uncomment the next line to test the router's internal secondary recovery:
		// panic("Oops, the recovery handler panicked too!")

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Custom Internal Server Error: Don't worry, we caught the panic!"))
	})

	// A route that intentionally triggers a panic to demonstrate the recovery mechanism.
	r.HandleFunc("/", "ANY", func(writer http.ResponseWriter, request *http.Request, context *router.Context) {
		panic("Something went terribly wrong")
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
