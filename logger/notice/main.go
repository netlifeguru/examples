package main

import "github.com/netlifeguru/logger"

func main() {
	closer, err := logger.Init(logger.Config{
		Dir:            "./logs",
		TerminalOutput: true,
	})
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	logger.Notice("important event", "order_id", 1)
}
