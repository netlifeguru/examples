package main

import (
	"os"

	"github.com/netlifeguru/db"
	scylla "github.com/netlifeguru/db-scylla"
)

func connectDB() (db.Conn, error) {
	conn := scylla.New()

	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if err := conn.CreatePool(cfg); err != nil {
		return nil, err
	}

	return conn.Fork(), nil
}
