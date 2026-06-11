package main

import (
	"log/slog"

	"github.com/netlifeguru/logger"
)

func main() {
	closer, err := logger.Init(logger.Config{
		Dir:            "./logs",
		TerminalOutput: true,
	})
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	slog.Info("user logged in",
		"user_id", 123,
		"role", "admin",
	)

	slog.Info("order processed",
		slog.Group("order",
			"id", 999,
			"amount", 49.99,
		),
		slog.Group("user",
			"id", 123,
			"email", "user@example.com",
		),
	)

	log := logger.With("service", "checkout")

	log.Info("payment completed",
		"order_id", 999,
		"status", "success",
	)
}
