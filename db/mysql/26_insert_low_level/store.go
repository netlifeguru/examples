package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const insertUserQuery = `
	INSERT INTO users (name, email, active)
	VALUES (?, ?, ?)
`

func InsertUser(ctx context.Context, conn db.Conn, name string, email string, active bool) (db.Result, error) {
	q, err := db.Raw(insertUserQuery, name, email, active)

	if err != nil {
		return nil, err
	}

	return conn.ExecCtx(ctx, q)
}
