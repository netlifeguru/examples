package main

import (
	"os"
	"strconv"
	"time"

	"github.com/netlifeguru/db"
	mysql "github.com/netlifeguru/db-mysql"
)

func connectDB() (db.Conn, error) {
	conn := mysql.New()

	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	cfg := db.Config{
		Identifier: "default",

		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,

		MaxConns:          50,
		MinConns:          5,
		MaxConnIdleTime:   10 * time.Minute,
		MaxConnLifetime:   2 * time.Hour,
		HealthCheckPeriod: 30 * time.Second,
		ConnectTimeout:    10 * time.Second,

		SSLMode:  "false",
		TimeZone: "Local",
	}

	if err := conn.CreatePool(cfg); err != nil {
		return nil, err
	}

	return conn.Fork(), nil
}
