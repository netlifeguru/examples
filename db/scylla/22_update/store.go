package main

import (
	"context"
	"fmt"

	"github.com/netlifeguru/db"
)

const updateUserByIDQuery = `
	UPDATE users_by_id
	SET name = ?, active = ?
	WHERE id = ?
`

func UpdateUser(ctx context.Context, conn db.Conn, id string, name string, active bool) error {
	if result, err := db.Update(ctx, conn, updateUserByIDQuery, name, active, id); err != nil {
		return err
	} else {
		fmt.Println(result)
	}

	return nil
}
