package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/ryudokung/Project-0/backend/internal/auth"
)

func main() {
	fmt.Println("Project-0 Backend Starting...")

	// Database Connection
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/project0?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize Auth Module
	authRepo := auth.NewRepository(db)
	authUseCase := auth.NewUseCase(authRepo, "your-very-secret-key") // In production, use env var
	authHandler := auth.NewHandler(authUseCase)

	// Routes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/api/v1/auth/login", authHandler.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
