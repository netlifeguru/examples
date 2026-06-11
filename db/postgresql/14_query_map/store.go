package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const selectUserMapsQuery = `
	SELECT *
	FROM users
	ORDER BY created_at DESC
	LIMIT $1
`

func ListUserMaps(ctx context.Context, conn db.Conn) ([]map[string]any, error) {
	q, err := db.Raw(selectUserMapsQuery, 10)
	if err != nil {
		return nil, err
	}

	return db.MapsQuery(ctx, conn, q)
}
