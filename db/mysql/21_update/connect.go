package main

import (
	"os"
	"strconv"

	"github.com/netlifeguru/db"
	mysql "github.com/netlifeguru/db-mysql"
)

func connectDB() (db.Conn, error) {
	conn := mysql.New()

	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
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
