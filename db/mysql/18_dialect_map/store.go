package main

import (
	"context"

	"github.com/netlifeguru/db"
)

func SelectUserMaps(ctx context.Context, conn db.Conn, queries Queries) ([]map[string]any, error) {
	return db.MapsDialect(ctx, conn, queries.ListUsers, 10)
}

func SelectUserMapsExp(ctx context.Context, conn db.Conn) ([]map[string]any, error) {
	return db.MapsDialect(ctx, conn, db.DialectSQL{
		Postgres: `SELECT * FROM users ORDER BY created_at DESC LIMIT $1`,
		Mysql:    `SELECT * FROM users ORDER BY created_at DESC LIMIT ?`,
	}, 10)
}
