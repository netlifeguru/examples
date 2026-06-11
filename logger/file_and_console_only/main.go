package main

import "github.com/netlifeguru/logger"

func main() {
	closer, _ := logger.Init(logger.Config{
		Dir:            "./logs",
		TerminalOutput: true,
	})
	defer closer.Close()

	logger.Notice("service is starting")
}
