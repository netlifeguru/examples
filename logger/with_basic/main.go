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

	log := logger.With(
		"service", "payment",
		"version", "v1",
	)

	log.Info("starting service")
	log.Error("failed to connect", "retry", true)
}
