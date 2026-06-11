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

	posts, err := ListPostMapsByUser(ctx, conn, queries, userID, 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("%v | user=%v | %v | published=%v | created_at=%v\n",
			post["id"],
			post["user_id"],
			post["title"],
			post["published"],
			post["created_at"],
		)
	}

	fmt.Println("")

	posts, err = ListPostMapsByUserExp(ctx, conn, queries, userID, 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("%v | user=%v | %v | published=%v | created_at=%v\n",
			post["id"],
			post["user_id"],
			post["title"],
			post["published"],
			post["created_at"],
		)
	}
}
