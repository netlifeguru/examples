package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, I'm using system env variables")
	}

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	events, err := getEvents(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Printf(
			"ID: %v | Type: %v | Payload: %v | Created: %v\n",
			event["id"],
			event["type"],
			event["payload"],
			event["created_at"],
		)
	}
}
