package main

import (
	"context"

	"github.com/netlifeguru/db"
)

func CountUsers(ctx context.Context, conn db.Conn, queries Queries) (int64, bool, error) {
	return db.ValueDialect[int64](ctx, conn, queries.CountUsers)
}

func CountUsersExp(ctx context.Context, conn db.Conn) (int64, bool, error) {
	return db.ValueDialect[int64](ctx, conn, db.DialectSQL{
		Postgres: `SELECT COUNT(*) FROM users`,
		Mysql:    `SELECT COUNT(*) FROM users`,
	})
}
