package main

import (
	"context"
	"github.com/netlifeguru/db"
)

const updateUserQuery = `
	UPDATE users
	SET name = ?, email = ?, active = ?
	WHERE id = ?
`

func UpdateUser(ctx context.Context, conn db.Conn, id int, name string, email string, active bool) (db.Result, error) {
	return db.Update(ctx, conn, updateUserQuery, name, email, active, id)
}
