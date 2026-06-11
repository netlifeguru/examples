package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/netlifeguru/db"
)

const insertUserIfNotExistsQuery = `
	INSERT INTO users_by_email (email, id, name, active, created_at)
	VALUES (?, ?, ?, ?, ?)
	IF NOT EXISTS
`

func InsertUserIfNotExists(ctx context.Context, conn db.Conn, name string, email string, active bool) (bool, string, error) {
	id := gocql.TimeUUID()
	createdAt := time.Now().UTC()

	rows, err := db.Maps(
		ctx,
		conn,
		insertUserIfNotExistsQuery,
		email,
		id,
		name,
		active,
		createdAt,
	)
	if err != nil {
		return false, "", err
	}

	if len(rows) == 0 {
		return false, "", nil
	}

	applied, ok := rows[0]["[applied]"].(bool)
	if !ok {
		return false, "", fmt.Errorf("missing [applied] value")
	}

	return applied, id.String(), nil
}
