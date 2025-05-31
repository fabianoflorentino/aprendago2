package main

import (
	"log"

	"github.com/fabianoflorentino/aprendago2/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
}
