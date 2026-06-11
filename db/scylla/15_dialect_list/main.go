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

	userID := "11111111-1111-1111-1111-111111111111"

	posts, err := ListPostsByUser(ctx, conn, queries, userID, 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("%s | user=%s | %s | published=%v | created_at=%s\n",
			post.ID,
			post.UserID,
			post.Title,
			post.Published,
			post.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}

	fmt.Println("")

	posts, err = ListPostsByUserExp(ctx, conn, queries, userID, 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("%s | user=%s | %s | published=%v | created_at=%s\n",
			post.ID,
			post.UserID,
			post.Title,
			post.Published,
			post.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}
