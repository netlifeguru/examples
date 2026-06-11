package main

import (
	"context"
	"github.com/netlifeguru/db"
)

func ListPostMapsByUser(ctx context.Context, conn db.Conn, queries Queries, userID string, limit int) ([]map[string]any, error) {
	return db.MapsDialect(ctx, conn, queries.ListPostsByUser, userID, limit)
}

func ListPostMapsByUserExp(ctx context.Context, conn db.Conn, queries Queries, userID string, limit int) ([]map[string]any, error) {
	return db.MapsDialect(ctx, conn, db.DialectSQL{
		Scylla: `SELECT * FROM posts_by_user WHERE user_id = ? LIMIT ?`,
	}, userID, limit)
}
