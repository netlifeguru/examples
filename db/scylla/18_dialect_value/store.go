package main

import (
	"context"
	"github.com/netlifeguru/db"
)

func CountPostsByUser(ctx context.Context, conn db.Conn, queries Queries, userID string) (int64, bool, error) {
	return db.ValueDialect[int64](ctx, conn, queries.CountPostsByUser, userID)
}

func CountPostsByUserExp(ctx context.Context, conn db.Conn, queries Queries, userID string) (int64, bool, error) {
	return db.ValueDialect[int64](ctx, conn, db.DialectSQL{
		Scylla: `SELECT COUNT(*) FROM posts_by_user WHERE user_id = ?`,
	}, userID)
}
