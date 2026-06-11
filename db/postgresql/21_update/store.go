package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const updateUserQuery = `
	UPDATE users
	SET name = $1, email = $2, active = $3
	WHERE id = $4
`

func UpdateUser(ctx context.Context, conn db.Conn, id int, name string, email string, active bool) (db.Result, error) {
	return db.Update(ctx, conn, updateUserQuery, name, email, active, id)
}
