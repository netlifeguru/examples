package main

import (
	"database/sql"
	"time"

	"github.com/netlifeguru/mapper"
)

type Event struct {
	ID        string         `db:"id"`
	Type      string         `db:"type"`
	Payload   map[string]any `db:"payload"`
	CreatedAt time.Time      `db:"created_at"`
}

func getEvents(db *sql.DB) ([]Event, error) {
	rows, err := db.Query(`
		SELECT id, type, payload, created_at
		FROM events
		ORDER BY created_at DESC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructSlice[Event](rows)
}
