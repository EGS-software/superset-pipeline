package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/EGS-software/superset-pipeline/etl-go/internal/model"
	"github.com/EGS-software/superset-pipeline/etl-go/internal/service"
	"github.com/go-resty/resty/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Error: DATABASE_URL environment variable is not set")
	}

	// --- RETRY LOGIC START ---
	var conn *pgx.Conn // Var used to try DB connect
	var err error
	ctx := context.Background()

	fmt.Println("⏳ Waiting for database to be ready...")
	for i := 1; i <= 10; i++ {
		conn, err = pgx.Connect(ctx, dbURL)
		if err == nil {
			break
		}
		fmt.Printf("Attempt %d/10: Database not ready yet, retrying in 2s...\n", i)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("❌ Could not connect to database after 10 attempts: %v", err)
	}

	defer conn.Close(ctx)
	// --- RETRY LOGIC END ---

	fmt.Println("✅ Successfully connected to Postgres!")

	// 3. Create table
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS pokemon_analytics (
       id INT PRIMARY KEY,
       name TEXT,
       generation INT,
       total_stats INT,
       type_1 TEXT,
       type_2 TEXT,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = conn.Exec(ctx, createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// 4. Init Collect
	client := resty.New()
	fmt.Println("🚀 Initializing Get Data...")

	for i := 1; i <= 151; i++ { // "i" will be used to get Pokémon by ID in API
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", i)
		resp, err := client.R().Get(url)

		if err != nil {
			fmt.Printf("❌ Error to get ID %d: %v\n", i, err)
			continue
		}

		var p model.PokemonAPI
		json.Unmarshal(resp.Body(), &p)

		dbData := service.TransformPokemon(p)

		// If data already exists, its replacement to new data with this rules
		upsertSQL := `
			INSERT INTO pokemon_analytics (id, name, generation, total_stats, type_1, type_2)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (id) DO UPDATE 
			SET name = EXCLUDED.name, 
			    total_stats = EXCLUDED.total_stats,
			    type_1 = EXCLUDED.type_1,
			    type_2 = EXCLUDED.type_2,
			    updated_at = NOW();
		`
		// We get the current context(db), sql string and arguments to replace data
		_, err = conn.Exec(
			ctx,
			upsertSQL,
			p.ID,
			p.Name,
			dbData.Generation,
			dbData.TotalStats,
			dbData.TypeOne,
			dbData.TypeTwo)
		if err != nil {
			fmt.Printf("❌ Error to save %s: %v\n", p.Name, err)
		} else {
			fmt.Printf("💾 [%d] %s salvo com BST: %d\n", p.ID, p.Name, dbData.TotalStats)
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("✨ Data collect completed!")
}
