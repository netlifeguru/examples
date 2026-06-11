package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const selectPostMapsByUserQuery = `
	SELECT *
	FROM posts_by_user
	WHERE user_id = ?
	LIMIT ?
`

func ListPostMapsByUser(ctx context.Context, conn db.Conn, userID string, limit int) ([]map[string]any, error) {
	q, err := db.Raw(selectPostMapsByUserQuery, userID, limit)
	if err != nil {
		return nil, err
	}

	return db.MapsQuery(ctx, conn, q)
}
