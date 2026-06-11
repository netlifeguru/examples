package main

import (
	"log/slog"

	"github.com/netlifeguru/logger"
)

func main() {
	closer, err := logger.Init(logger.Config{
		Dir:             "./logs",
		TerminalOutput:  true,
		DisableColors:   false,
		MinLevel:        slog.LevelInfo,
		ConsoleMinLevel: slog.LevelDebug,
		MaxFileSize:     10 * 1024 * 1024,
		MaxLogFiles:     5,
		AddSource:       true,
	})
	if err != nil {
		slog.Error(err.Error())
	}
	defer closer.Close()

	slog.Info("logger initialized")
}
