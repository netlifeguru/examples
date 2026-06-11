package main

import (
	"fmt"
	"log"
	"time"

	"github.com/netlifeguru/mapper"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Active    bool
	CreatedAt time.Time
}

func main() {
	row := mapper.Row{
		"id":         1,
		"name":       "John Doe",
		"email":      "john@example.com",
		"active":     true,
		"created_at": time.Date(2026, 5, 7, 11, 15, 21, 0, time.UTC),
	}

	user, err := scanUser(row)
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

func scanUser(row mapper.Row) (User, error) {
	id, ok := row.Int("id")
	if !ok {
		return User{}, fmt.Errorf("invalid id")
	}

	name, ok := row.String("name")
	if !ok {
		return User{}, fmt.Errorf("invalid name")
	}

	email, ok := row.String("email")
	if !ok {
		return User{}, fmt.Errorf("invalid email")
	}

	active, ok := row.Bool("active")
	if !ok {
		return User{}, fmt.Errorf("invalid active")
	}

	createdAt, ok := row.Time("created_at")
	if !ok {
		return User{}, fmt.Errorf("invalid created_at")
	}

	return User{
		ID:        id,
		Name:      name,
		Email:     email,
		Active:    active,
		CreatedAt: createdAt,
	}, nil
}
