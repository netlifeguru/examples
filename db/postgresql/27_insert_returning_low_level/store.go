package main

import (
	"context"
	"errors"

	"github.com/netlifeguru/db"
)

const insertUserQuery = `
	INSERT INTO users (name, email, active)
	VALUES ($1, $2, $3)
	RETURNING id
`

func InsertUser(ctx context.Context, conn db.Conn, name string, email string, active bool) (int64, error) {
	id, found, err := db.Value[int64](ctx, conn, insertUserQuery, name, email, active)
	if err != nil {
		return 0, err
	}

	if !found {
		return 0, errors.New("insert did not return id")
	}

	return id, nil
}
