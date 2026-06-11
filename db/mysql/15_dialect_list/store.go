package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func ListUsers(ctx context.Context, conn db.Conn, queries Queries) ([]User, error) {
	return db.ListDialect[User](ctx, conn, queries.ListUsers, 10)
}

func ListUsersExp(ctx context.Context, conn db.Conn) ([]User, error) {
	return db.ListDialect[User](ctx, conn, db.DialectSQL{
		Postgres: `SELECT * FROM users ORDER BY created_at DESC LIMIT $1`,
		Mysql:    `SELECT * FROM users ORDER BY created_at DESC LIMIT ?`,
	}, 10)
}
