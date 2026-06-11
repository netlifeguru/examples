package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, I'm using system env variables")
	}

	conn, err := connectDB()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connected to %s", conn.DriverName())
}
