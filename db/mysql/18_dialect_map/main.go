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

	users, err := SelectUserMaps(ctx, conn, queries)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("%v | %v | %v | active=%v | created_at=%v\n",
			user["id"],
			user["name"],
			user["email"],
			user["active"],
			user["created_at"],
		)
	}

	fmt.Println("")

	users, err = SelectUserMapsExp(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("%v | %v | %v | active=%v | created_at=%v\n",
			user["id"],
			user["name"],
			user["email"],
			user["active"],
			user["created_at"],
		)
	}
}
