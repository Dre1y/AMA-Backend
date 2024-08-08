package main

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Prepare the tern command with its arguments
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// Capture the standard output and error
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running tern migrate: %v\n%s", err, stderr.String())
	}

	log.Println("Migration executed successfully.")
	log.Printf("Command output: %s", out.String())
}
