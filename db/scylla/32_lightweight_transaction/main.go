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

	applied, id, err := InsertUserIfNotExists(ctx, conn, "LWT User", "lwt.user@example.com", true)
	if err != nil {
		log.Fatal(err)
	}

	if !applied {
		fmt.Printf("user already exists: %s\n", id)
		return
	}

	fmt.Printf("inserted user id=%s\n", id)
}
