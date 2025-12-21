package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/ryudokung/Project-0/backend/internal/auth"
	"github.com/ryudokung/Project-0/backend/internal/mech"
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
	jwtSecret := os.Getenv("PRIVY_APP_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-change-me"
	}
	authRepo := auth.NewRepository(db)
	authUseCase := auth.NewUseCase(authRepo, jwtSecret)
	authHandler := auth.NewHandler(authUseCase)

	// Initialize Mech Module
	mechRepo := mech.NewRepository(db)
	mechUseCase := mech.NewUseCase(mechRepo)
	mechHandler := mech.NewHandler(mechUseCase)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/mechs/mint-starter", mechHandler.MintStarter)
	mux.HandleFunc("/api/v1/mechs", mechHandler.ListMechs)

	// Simple CORS Middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		mux.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
