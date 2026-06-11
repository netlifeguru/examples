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

	connections, err := connectDB()

	if err != nil {
		log.Fatal(err)
	}

	mainDB := connections["main"]
	analyticsDB := connections["analytics"]

	log.Printf("main connected to %s", mainDB.DriverName())
	log.Printf("analytics connected to %s", analyticsDB.DriverName())
}
