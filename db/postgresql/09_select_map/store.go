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

func SelectUserMaps(ctx context.Context, conn db.Conn) ([]map[string]any, error) {
	return db.Maps(ctx, conn, selectUserMapsQuery, 10)
}
