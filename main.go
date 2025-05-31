package main

import (
	"log"

	"github.com/fabianoflorentino/aprendago2/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
