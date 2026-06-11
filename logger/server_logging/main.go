package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/netlifeguru/logger"
)

func main() {
	config := logger.Config{
		Dir:             "./logs",
		TerminalOutput:  true,
		DisableColors:   false,
		MinLevel:        slog.LevelInfo,
		ConsoleMinLevel: slog.LevelDebug,
		MaxFileSize:     1 * 1024 * 1024,
		MaxLogFiles:     5,
		AddSource:       true,
	}

	closer, err := logger.Init(config)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log := logger.With(
			"service", "api",
		).With(
			"method", r.Method,
			"path", r.URL.Path,
		)

		log.Info("request started")

		log = log.With("request_id", "req-123")

		log.Info("processing")

		log.Error("failed", "error", "something broke")

		reqLog := logger.With(
			"request_id", r.Header.Get("X-Request-Id"),
			"path", r.URL.Path,
		)

		reqLog.Info("request start")
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("server is starting", "port", 8080)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server crashed", "error", err)
		os.Exit(1)
	}
}
