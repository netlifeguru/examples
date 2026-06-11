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

	log := logger.With("service", "checkout")

	reqLog := log.With("request_id", "xyz")

	reqLog.InfoContext(ctx, "starting checkout")
}
