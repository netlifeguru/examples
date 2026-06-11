package main

import (
	"context"

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

	logger.NoticeContext(ctx, "important event", "order_id", 1)
}
