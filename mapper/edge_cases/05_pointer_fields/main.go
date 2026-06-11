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

	profiles, err := getProfiles(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, profile := range profiles {
		fmt.Printf(
			"ID: %s | Name: %s | Bio: %s | Website: %s\n",
			profile.ID,
			profile.Name,
			stringPtr(profile.Bio),
			stringPtr(profile.Website),
		)
	}
}

func stringPtr(value *string) string {
	if value == nil {
		return "NULL"
	}

	return *value
}
