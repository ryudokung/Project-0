package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbURL := "postgres://user:password@localhost:5432/project0?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Drop all tables and types
	fmt.Println("Dropping existing tables and types...")
	dropQuery := `
		DROP TABLE IF EXISTS exploration_nodes CASCADE;
		DROP TABLE IF EXISTS exploration_sessions CASCADE;
		DROP TABLE IF EXISTS expeditions CASCADE;
		DROP TABLE IF EXISTS nodes CASCADE;
		DROP TABLE IF EXISTS sub_sectors CASCADE;
		DROP TABLE IF EXISTS sectors CASCADE;
		DROP TABLE IF EXISTS planet_locations CASCADE;
		DROP TABLE IF EXISTS stars CASCADE;
		DROP TABLE IF EXISTS pilot_stats CASCADE;
		DROP TABLE IF EXISTS parts CASCADE;
		DROP TABLE IF EXISTS items CASCADE;
		DROP TABLE IF EXISTS vehicles CASCADE;
		DROP TABLE IF EXISTS characters CASCADE;
		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS gacha_stats CASCADE;
		DROP TYPE IF EXISTS vehicle_type CASCADE;
		DROP TYPE IF EXISTS vehicle_class CASCADE;
		DROP TYPE IF EXISTS rarity_tier CASCADE;
		DROP TYPE IF EXISTS saga_status CASCADE;
		DROP TYPE IF EXISTS vehicle_status CASCADE;
		DROP TYPE IF EXISTS item_type CASCADE;
		DROP TYPE IF EXISTS item_condition CASCADE;
		DROP TYPE IF EXISTS node_type CASCADE;
		DROP FUNCTION IF EXISTS update_updated_at_column CASCADE;
	`
	_, err = db.Exec(dropQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Read init.sql
	content, err := os.ReadFile("init.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Execute init.sql
	fmt.Println("Executing init.sql...")
	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database re-initialized successfully!")
}
