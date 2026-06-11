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
	q, err := db.Raw(countPostsByUserQuery, userID)
	if err != nil {
		return 0, false, err
	}

	return db.ValueQuery[int64](ctx, conn, q)
}
