package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Estruturas para mapear a PokeAPI

func main() {
	// 1. Get .env configs
	_ = godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not configured")
	}

	// 2. Connect in Postgres
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("Error to connect in DB: %v", err)
	}
	defer conn.Close(ctx)

	fmt.Println("✅ Connected in Postgres!")
}
