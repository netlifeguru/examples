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

	id := 1

	user, err := SelectUser(ctx, conn, queries, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d | %s | %s | active=%v | created_at=%s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)

	user, err = SelectUserExp(ctx, conn, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d | %s | %s | active=%v | created_at=%s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
