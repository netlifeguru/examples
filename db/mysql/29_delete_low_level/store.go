package main

import (
	"context"

	"github.com/netlifeguru/db"
)

const deleteUserQuery = `DELETE FROM users WHERE id = ?`

func DeleteUser(ctx context.Context, conn db.Conn, id int) (db.Result, error) {

	result, err := db.Delete(ctx, conn, deleteUserQuery, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}
