package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const insertUserQuery = `
	INSERT INTO users (name, email, active)
	VALUES ($1, $2, $3)
`

func InsertUser(ctx context.Context, conn db.Conn, name string, email string, active bool) (db.Result, error) {
	result, err := db.Insert(ctx, conn, insertUserQuery, name, email, active)

	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 1 {

	}

	return result, nil
}
