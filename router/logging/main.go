package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/logger"
	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// --- Logger Configuration ---
	// The logger is a high-performance external package that handles
	// asynchronous file writing and log rotation.
	config := logger.Config{
		Dir:             "./logs",          // Directory for log storage
		TerminalOutput:  false,             // Production-ready: set to false to disable console coloring
		DisableColors:   false,             // disabling color highlights
		MinLevel:        slog.LevelInfo,    // Minimum log level to capture
		ConsoleMinLevel: slog.LevelDebug,   // Console minimum log level to capture
		MaxFileSize:     100 * 1024 * 1024, // 100MB per file before rotation
		MaxLogFiles:     10,                // Retention policy (number of old log files)
	}

	// Initialize the global default slog handler
	closer, err := logger.Init(config)
	if err != nil {
		slog.Error("Failed to start logger", "error", err)
		os.Exit(1)
	}

	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			slog.Error("Failed to close logger", "error", err)
		}
	}(closer)

	// --- Middleware Registration ---
	// should be completely removed to eliminate per-request overhead.
	r.Use(router.Logger())

	r.GET("/", func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`Hello World`))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
