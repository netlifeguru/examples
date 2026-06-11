package main

import (
	"database/sql"

	"github.com/netlifeguru/mapper"
)

type UserProfile struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	Bio       sql.NullString `db:"bio"`
	Website   sql.NullString `db:"website"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

func getProfiles(db *sql.DB) ([]UserProfile, error) {
	rows, err := db.Query(`
		SELECT
			u.id,
			u.name,
			up.bio,
			up.website,
			up.updated_at
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

func nullString(value sql.NullString) string {
	if !value.Valid {
		return "NULL"
	}

	return value.String
}

func nullTime(value sql.NullTime) string {
	if !value.Valid {
		return "NULL"
	}

	return value.Time.Format("2006-01-02 15:04:05")
}
