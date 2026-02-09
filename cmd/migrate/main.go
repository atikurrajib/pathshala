package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Database Connected. Starting Migration...")

	files, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(files)

	for _, file := range files {
		fmt.Printf("Running %s... ", file)

		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		_, err = conn.Exec(context.Background(), string(content))
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Println("SKIPPED (Already exists)")
				continue
			}
			log.Fatalf("\nError running migration %s: %v", file, err)
		}

		fmt.Println("DONE!")
	}

	fmt.Println("\nAll migrations executed successfully!")
}

