package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const countPostsByUserQuery = `
	SELECT COUNT(*)
	FROM posts_by_user
	WHERE user_id = ?
`

func CountPostsByUser(ctx context.Context, conn db.Conn, userID string) (int64, bool, error) {
	return db.Value[int64](ctx, conn, countPostsByUserQuery, userID)
}
