package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/netlifeguru/db"
	postgres "github.com/netlifeguru/db-postgres"
)

func connectDB() (map[string]db.Conn, error) {
	base := postgres.New()

	configs, err := loadPoolConfigs()
	if err != nil {
		return nil, err
	}

	connections := make(map[string]db.Conn, len(configs))

	for _, cfg := range configs {
		if err := base.CreatePool(cfg); err != nil {
			return nil, fmt.Errorf("create pool %s: %w", cfg.Identifier, err)
		}

		connections[cfg.Identifier] = base.Fork()
	}

	return connections, nil
}

func loadPoolConfigs() ([]db.Config, error) {
	first, err := loadPoolConfig("_1")
	if err != nil {
		return nil, err
	}

	second, err := loadPoolConfig("_2")
	if err != nil {
		return nil, err
	}

	return []db.Config{first, second}, nil
}

func loadPoolConfig(suffix string) (db.Config, error) {
	cfg := db.Config{
		Identifier: os.Getenv("DB_IDENTIFIER" + suffix),

		Host:     os.Getenv("DB_HOST" + suffix),
		Database: os.Getenv("DB_NAME" + suffix),
		Username: os.Getenv("DB_USER" + suffix),
		Password: os.Getenv("DB_PASSWORD" + suffix),

		MaxConns:          50,
		MinConns:          5,
		MaxConnIdleTime:   10 * time.Minute,
		MaxConnLifetime:   2 * time.Hour,
		HealthCheckPeriod: 30 * time.Second,
		ConnectTimeout:    10 * time.Second,

		SSLMode:  "disable",
		TimeZone: "UTC",
	}

	if port := os.Getenv("DB_PORT" + suffix); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil {
			return db.Config{}, fmt.Errorf("invalid DB_PORT%s: %w", suffix, err)
		}
		cfg.Port = n
	}

	return cfg, nil
}
