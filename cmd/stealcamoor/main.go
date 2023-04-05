package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/0xmichalis/stealcamoor/pkg/stealcamoor"
)

func main() {
	err := godotenv.Load()
	if err == nil {
		log.Printf("Loading config from .env file")
	} else {
		log.Printf("No .env file found, will read config from environment variables")
	}

	l, err := stealcamoor.New()
	if err != nil {
		log.Fatalf("Failed to instantiate stealcamoor: %v", err)
	}

	l.Start()
}
