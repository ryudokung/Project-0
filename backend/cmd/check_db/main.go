package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := "postgres://user:password@localhost:5432/project0?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Latest 5 nodes:")
	rows, err = db.Query("SELECT id, expedition_id, name, type, choices FROM nodes ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, expID, name, nodeType, choices string
		if err := rows.Scan(&id, &expID, &name, &nodeType, &choices); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, ExpID: %s, Name: %s, Type: %s, Choices: %s\n", id, expID, name, nodeType, choices)
	}
}
