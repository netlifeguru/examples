package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type UserProfile struct {
	ID      string  `db:"id"`
	Name    string  `db:"name"`
	Bio     *string `db:"bio"`
	Website *string `db:"website"`
}

func getProfiles(db *sql.DB) ([]UserProfile, error) {
	rows, err := db.Query(`
		SELECT
			u.id,
			u.name,
			up.bio,
			up.website
		FROM users u
		LEFT JOIN user_profiles up ON up.user_id = u.id
		ORDER BY u.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return mapper.ScanStructSlice[UserProfile](rows)
}
