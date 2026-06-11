package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// --- ENABLE PROFILING ---
	// This starts a dedicated internal HTTP server for profiling.
	// You can access the dashboard at: http://localhost:6060/debug/pprof/
	// Use this to analyze performance, memory leaks, and CPU spikes.
	r.EnableProfiling("localhost:6060")

	r.GET("/", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Write([]byte("Server is running with profiling enabled."))
	})

	// Main application server
	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
