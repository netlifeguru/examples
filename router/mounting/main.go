package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {

	r := router.New()

	// --- Interoperability with net/http ---
	// These methods allow you to mount standard Go handlers that do not use
	// your custom *router.Context. This is crucial for integrating third-party
	// libraries (e.g., Prometheus metrics, Swagger UI, pprof) or legacy code.

	// 1. Mount: Attaching a standard http.Handler interface
	// Use this when you have a struct that implements ServeHTTP or an external library handler.
	r.Mount("/service", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mounted standard service (http.Handler)"))
	}))

	// 2. MountFunc: Attaching a standard function
	// A syntactic sugar wrapper for Mount. It accepts a standard function signature
	// func(w, r) directly, saving you the explicit cast to http.HandlerFunc.
	r.MountFunc("/serviceFunc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mounted standard function")
	})

	// 3. Handling Parameters in Standard Handlers
	// Even though these are standard handlers, the router still parses the path parameters.
	// You can access them using the standard Go 1.22+ method: r.PathValue("key").
	r.Mount("/service/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the parameter natively from the request object
		id := r.PathValue("id")
		fmt.Fprintf(w, "Mounted service with ID: %s", id)
	}))

	r.MountFunc("/serviceFunc/{id}/abc", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Mounted func with ID: %s", id)
	})

	// 4. Wildcards & Catch-All Routes
	// Wildcards allow you to forward an entire URL branch to a standard handler.
	// This is commonly used for serving static files, SPAs, or mounting sub-routers.
	r.Mount("/services/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Matches everything starting with /service/
		fmt.Fprintf(w, "Wildcard captured path: %s", r.URL.Path)
	}))

	r.MountFunc("/servicesFunc/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Wildcard func captured path: %s", r.URL.Path)
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
