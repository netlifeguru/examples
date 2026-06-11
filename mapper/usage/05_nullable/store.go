package main

import (
	"database/sql"
	"time"

	"github.com/netlifeguru/mapper"
)

type UserProfile struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	Email     string         `db:"email"`
	Bio       sql.NullString `db:"bio"`
	AvatarURL sql.NullString `db:"avatar_url"`
	Website   sql.NullString `db:"website"`
	Birthday  sql.NullTime   `db:"birthday"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	CreatedAt time.Time      `db:"created_at"`
}

func getUserProfiles(db *sql.DB) ([]UserProfile, error) {
	rows, err := db.Query(`
		SELECT
			u.id,
			u.name,
			u.email,
			up.bio,
			up.avatar_url,
			up.website,
			up.birthday,
			up.updated_at,
			u.created_at
		FROM users u
		LEFT JOIN user_profiles up ON up.user_id = u.id
		ORDER BY u.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []UserProfile

	err = mapper.ScanStructRows[UserProfile](rows, func(profile *UserProfile) error {
		profiles = append(profiles, *profile)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return profiles, nil
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
