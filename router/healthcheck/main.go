package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// Liveness Probe (Default: /healthz)
	// Returns 200 OK as long as the process is running.
	// If this endpoint fails, the orchestrator (e.g., K8s) will restart the pod.
	// You can pass an empty string "" to use the default path "/healthz".
	r.Liveness("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Readiness Probe (Default: /readyz)
	// Returns 200 OK only when the application is ready to accept traffic.
	// The router automatically toggles r.IsReady() to false when it receives
	// a shutdown signal (SIGTERM/SIGINT), ensuring zero-downtime deployments.

	r.Readiness("/readyz", func(w http.ResponseWriter, req *http.Request) {
		// Check the router's internal state
		if r.IsReady() {
			// You can extend this to check DB connections, Redis, etc.
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("READY"))
		} else {
			// Return 503 Service Unavailable during startup or shutdown
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("NOT_READY"))
		}
	})

	// Standard API route
	r.GET("/", func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		w.Write([]byte("Welcome to NetLifeGuru Router"))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
