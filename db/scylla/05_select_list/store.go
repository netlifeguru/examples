package main

import (
	"context"
	"time"

	"github.com/netlifeguru/db"
)

type Post struct {
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	Published bool      `db:"published"`
}

const selectPostsByUserQuery = `
	SELECT *
	FROM posts_by_user
	WHERE user_id = ?
	LIMIT ?
`

func ListPostsByUser(ctx context.Context, conn db.Conn, userID string, limit int) ([]Post, error) {
	return db.List[Post](ctx, conn, selectPostsByUserQuery, userID, limit)
}
