package main

import (
	"log/slog"

	"github.com/netlifeguru/logger"
)

func main() {
	closer, err := logger.Init(logger.Config{
		TerminalOutput:  true,
		ConsoleMinLevel: slog.LevelDebug,
	})

	if err != nil {
		slog.Error(err.Error())
	}

	defer closer.Close()

	slog.Debug("debug message")
	slog.Info("application started")
	slog.Warn("warning message")
	slog.Error("error message")

	slog.Info("user authenticated",
		slog.Int("user_id", 42),
		slog.String("role", "admin"),
	)
}
