package main

import (
	"context"
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

	ctx := context.Background()

	slog.InfoContext(ctx, "request received",
		"user_id", 123,
		"endpoint", "/api/pay",
	)

	slog.ErrorContext(ctx, "processing failed",
		"error", "timeout",
	)

	logger.NoticeContext(ctx, "important event",
		"order_id", 999,
	)
}
