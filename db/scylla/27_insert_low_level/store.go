package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"

	"github.com/gocql/gocql"
)

const insertUserByIDQuery = `
	INSERT INTO users_by_id (id, email, name, active, created_at)
	VALUES (?, ?, ?, ?, ?)
`

const insertUserByEmailQuery = `
	INSERT INTO users_by_email (email, id, name, active, created_at)
	VALUES (?, ?, ?, ?, ?)
`

func InsertUser(ctx context.Context, conn db.Conn, name string, email string, active bool) (string, error) {
	id := gocql.TimeUUID()
	createdAt := time.Now().UTC()

	q, err := db.Raw(insertUserByIDQuery, id, email, name, active, createdAt)
	if err != nil {
		return "", err
	}

	if _, err := conn.ExecCtx(ctx, q); err != nil {
		return "", err
	}

	q, err = db.Raw(insertUserByEmailQuery, email, id, name, active, createdAt)
	if err != nil {
		return "", err
	}

	if _, err := conn.ExecCtx(ctx, q); err != nil {
		return "", err
	}

	return id.String(), nil
}
