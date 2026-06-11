package main

import (
	"log/slog"

	"github.com/netlifeguru/logger"
)

func main() {
	closer, err := logger.Init(logger.Config{
		//TerminalOutput: true,
	})

	if err != nil {
		slog.Error(err.Error())
	}

	defer closer.Close()

	slog.Info("hello world")
}
