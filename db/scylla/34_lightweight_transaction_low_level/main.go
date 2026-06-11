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

	applied, id, err := InsertUserIfNotExists(ctx, conn, "LWT Low Level User", "lwt.low.level@example.com", true)
	if err != nil {
		log.Fatal(err)
	}

	if !applied {
		fmt.Println("user already exists")
		return
	}

	fmt.Printf("inserted user id=%s\n", id)
}
