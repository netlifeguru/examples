package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, I'm using system env variables")
	}

	conn, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	queries, err := LoadQueries(conn)
	if err != nil {
		log.Fatal(err)
	}

	email := "jane@example.com"

	user, found, err := GetUserByEmail(ctx, conn, queries, email)
	if err != nil {
		log.Fatal(err)
	}

	if !found {
		log.Printf("user with email %q was not found", email)
		return
	}

	fmt.Printf("%s | %s | %s | active=%v | created_at=%s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)

	user, found, err = GetUserByEmailExp(ctx, conn, queries, email)
	if err != nil {
		log.Fatal(err)
	}

	if !found {
		log.Printf("user with email %q was not found", email)
		return
	}

	fmt.Printf("%s | %s | %s | active=%v | created_at=%s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
