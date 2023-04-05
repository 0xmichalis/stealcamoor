package main

import (
	"log"

	"github.com/0xmichalis/stealcamoor/pkg/stealcamoor"
)

func main() {
	l, err := stealcamoor.New()
	if err != nil {
		log.Fatalf("Failed to instantiate stealcamoor: %v", err)
	}

	l.Start()
}