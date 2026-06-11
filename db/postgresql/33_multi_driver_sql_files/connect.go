package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/netlifeguru/db"
	mysql "github.com/netlifeguru/db-mysql"
	postgres "github.com/netlifeguru/db-postgres"
)

func connectDB() (map[string]db.Conn, error) {
	connections := make(map[string]db.Conn)

	mysqlConn, err := connectMySQL()
	if err != nil {
		return nil, err
	}

	postgresConn, err := connectPostgreSQL()
	if err != nil {
		return nil, err
	}

	connections["mysql"] = mysqlConn
	connections["postgresql"] = postgresConn

	return connections, nil
}

func connectMySQL() (db.Conn, error) {
	conn := mysql.New()

	cfg := db.Config{
		Identifier: "mysql",

		Host:     os.Getenv("MYSQL_DB_HOST"),
		Database: os.Getenv("MYSQL_DB_NAME"),
		Username: os.Getenv("MYSQL_DB_USER"),
		Password: os.Getenv("MYSQL_DB_PASSWORD"),
	}

	if port := os.Getenv("MYSQL_DB_PORT"); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("invalid MYSQL_DB_PORT: %w", err)
		}
		cfg.Port = n
	}

	if err := conn.CreatePool(cfg); err != nil {
		return nil, err
	}

	return conn.Fork(), nil
}

func connectPostgreSQL() (db.Conn, error) {
	conn := postgres.New()

	cfg := db.Config{
		Identifier: "postgresql",

		Host:     os.Getenv("POSTGRES_DB_HOST"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		Username: os.Getenv("POSTGRES_DB_USER"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
		SSLMode:  "disable",
		TimeZone: "UTC",
	}

	if port := os.Getenv("POSTGRES_DB_PORT"); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("invalid POSTGRES_DB_PORT: %w", err)
		}
		cfg.Port = n
	}

	if err := conn.CreatePool(cfg); err != nil {
		return nil, err
	}

	return conn.Fork(), nil
}
