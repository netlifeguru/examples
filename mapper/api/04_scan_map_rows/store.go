package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

func getEvents(db *sql.DB) ([]map[string]any, error) {
	rows, err := db.Query(`
		SELECT id, type, payload, created_at
		FROM events
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []map[string]any

	err = mapper.ScanMapRows(rows, func(row map[string]any) error {
		events = append(events, row)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return events, nil
}
