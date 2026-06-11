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

	id := 1

	user, found, err := SelectUser(ctx, conn, id)
	if err != nil {
		log.Fatal(err)
	}

	if !found {
		log.Printf("user with id %d was not found", id)
		return
	}

	fmt.Printf("%d | %s | %s | active=%v | created_at=%s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Active,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}
