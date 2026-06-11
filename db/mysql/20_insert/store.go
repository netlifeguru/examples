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
	return db.Insert(ctx, conn, insertUserQuery, name, email, active)
}
