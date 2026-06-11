package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const countUsersQuery = `
	SELECT COUNT(*)
	FROM users
`

func CountUsers(ctx context.Context, conn db.Conn) (int64, bool, error) {
	return db.Value[int64](ctx, conn, countUsersQuery)
}
