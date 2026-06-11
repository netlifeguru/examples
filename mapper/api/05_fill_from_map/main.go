package main

import (
	"fmt"
	"log"
	"time"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	row := map[string]any{
		"id":         1,
		"name":       "John Doe",
		"email":      "john@example.com",
		"active":     true,
		"created_at": time.Date(2026, 5, 7, 11, 15, 21, 0, time.UTC),
	}

	var user User

	err := mapper.FillFromMap(&user, row)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"ID: %d | Name: %s | Email: %s | Active: %t | Created: %s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
