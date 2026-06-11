package main

import (
	"os"
	"strconv"
	"time"

	"github.com/netlifeguru/db"
	scylla "github.com/netlifeguru/db-scylla"
)

func connectDB() (db.Conn, error) {
	conn := scylla.New()

	cfg := db.Config{
		Identifier: "default",

		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"), // Scylla keyspace
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),

		MaxConns:          50,
		MinConns:          5,
		MaxConnIdleTime:   10 * time.Minute,
		MaxConnLifetime:   2 * time.Hour,
		HealthCheckPeriod: 30 * time.Second,
		ConnectTimeout:    10 * time.Second,

		Consistency: "local_quorum",
	}

	if port := os.Getenv("DB_PORT"); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil {
			return nil, err
		}
		cfg.Port = n
	}

	if err := conn.CreatePool(cfg); err != nil {
		return nil, err
	}

	return conn.Fork(), nil
}
