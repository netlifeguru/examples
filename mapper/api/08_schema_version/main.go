package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/netlifeguru/mapper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, I'm using system env variables")
	}

	// Set this to your current migration/schema version.
	// When the schema changes, update this value.
	mapper.SetSchemaVersion("2026_05_15_001")

	fmt.Println("Schema version:", mapper.CurrentSchemaVersion())

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf(
			"ID: %d | Name: %s | Email: %s | Active: %t | Created: %s\n",
			user.ID,
			user.Name,
			user.Email,
			user.Active,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}
