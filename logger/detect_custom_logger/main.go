package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/netlifeguru/logger"
)

func isLoggerEnabled() bool {
	return slog.Default().Enabled(context.Background(), logger.LevelNotice)
}

func main() {
	closer, err := logger.Init(logger.Config{
		Dir:            "./logs",
		TerminalOutput: true,
	})
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	enabled := isLoggerEnabled()

	fmt.Println(enabled)
}
