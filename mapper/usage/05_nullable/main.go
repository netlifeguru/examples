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

	profiles, err := getUserProfiles(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, profile := range profiles {
		fmt.Printf(
			"User: %s <%s> | Bio: %s | Website: %s | Updated: %s\n",
			profile.Name,
			profile.Email,
			nullString(profile.Bio),
			nullString(profile.Website),
			nullTime(profile.UpdatedAt),
		)
	}
}
