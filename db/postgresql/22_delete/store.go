package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const deleteUserQuery = `DELETE FROM users WHERE id = $1`

func DeleteUser(ctx context.Context, conn db.Conn, id int) (db.Result, error) {
	return db.Delete(ctx, conn, deleteUserQuery, id)
}
