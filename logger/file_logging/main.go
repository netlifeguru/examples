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
	cfg := logger.Config{
		Dir:             "./logs",
		TerminalOutput:  true,
		DisableColors:   false,
		MinLevel:        slog.LevelInfo,
		ConsoleMinLevel: slog.LevelDebug,
		MaxFileSize:     1 * 1024 * 1024,
		MaxLogFiles:     5,
		AddSource:       true,
	}

	closer, err := logger.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("hello world")
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
