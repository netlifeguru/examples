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

	id := "11111111-1111-1111-1111-111111111111"

	if err := UpdateUser(ctx, conn, id, "Updated Scylla User", true); err != nil {
		log.Fatal(err)
	}

	fmt.Println("user updated")
}
