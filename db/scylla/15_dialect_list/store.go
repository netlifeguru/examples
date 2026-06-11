package main

import (
	"context"
	"github.com/netlifeguru/db"
	"time"
)

type Post struct {
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	Published bool      `db:"published"`
}

func ListPostsByUser(ctx context.Context, conn db.Conn, queries Queries, userID string, limit int) ([]Post, error) {
	return db.ListDialect[Post](ctx, conn, queries.ListPostsByUser, userID, limit)
}

func ListPostsByUserExp(ctx context.Context, conn db.Conn, queries Queries, userID string, limit int) ([]Post, error) {
	return db.ListDialect[Post](ctx, conn, db.DialectSQL{
		Scylla: `SELECT * FROM posts_by_user WHERE user_id = ? LIMIT ?`,
	}, userID, limit)
}
