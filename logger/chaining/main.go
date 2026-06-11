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

	base := logger.With("service", "auth")

	reqLogger := base.With(
		"request_id", "abc-123",
		"user_id", 42,
	)

	reqLogger.Info("login attempt")

	dbLogger := reqLogger.With("component", "db")

	dbLogger.Debug("query executed", "sql", "SELECT * FROM users")
}
