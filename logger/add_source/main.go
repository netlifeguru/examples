package main

import (
	"log/slog"

	"github.com/netlifeguru/logger"
)

func main() {
	closer, err := logger.Init(logger.Config{
		Dir:            "./logs",
		TerminalOutput: true,
		AddSource:      true,
	})
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	slog.Warn("important event", "order_id", 1)
}
