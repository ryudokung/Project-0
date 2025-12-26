package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/ryudokung/Project-0/backend/internal/auth"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/combat"
	"github.com/ryudokung/Project-0/backend/internal/exploration"
	"github.com/ryudokung/Project-0/backend/internal/gacha"
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

	// Initialize Game/Pilot Module
	gameRepo := game.NewRepository(db)
	vehicleRepo := vehicle.NewRepository(db) // Move up to use in gameUseCase
	gameUseCase := game.NewUseCase(gameRepo, vehicleRepo)

	// Initialize Auth Module
	jwtSecret := os.Getenv("PRIVY_APP_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-change-me"
	}
	authRepo := auth.NewRepository(db)
	authUseCase := auth.NewUseCase(authRepo, gameUseCase, jwtSecret)
	authHandler := auth.NewHandler(authUseCase)

	// Initialize Vehicle Module
	vehicleUseCase := vehicle.NewUseCase(vehicleRepo)
	vehicleHandler := vehicle.NewHandler(vehicleUseCase)

	// Initialize Combat Module
	combatEngine := combat.NewEngine()
	combatService := combat.NewService(combatEngine)
	combatHandler := combat.NewHandler(combatService, vehicleRepo, gameRepo)

	// Initialize Exploration Module
	explorationRepo := exploration.NewRepository(db)
	explorationService := exploration.NewService(explorationRepo, vehicleUseCase, gameRepo)
	explorationHandler := exploration.NewHandler(explorationService)

	// Initialize Game Handler
	gameHandler := game.NewHandler(gameUseCase, gameRepo)

	// Initialize Gacha Module
	gachaRepo := gacha.NewRepository(db)
	gachaUC := gacha.NewUseCase(gachaRepo, gameRepo, vehicleRepo)
	gachaHandler := gacha.NewHandler(gachaUC)

	// Routes
	mux := http.NewServeMux()
	
	// Public Routes
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/auth/signup", authHandler.Signup)
	mux.HandleFunc("/api/v1/exploration/universe-map", explorationHandler.GetUniverseMap)
	
	// Protected Routes Middleware
	authMiddleware := auth.Middleware(authUseCase)

	// Protected Routes
	mux.Handle("/api/v1/auth/me", authMiddleware(http.HandlerFunc(authHandler.GetMe)))
	mux.Handle("/api/v1/auth/link-wallet", authMiddleware(http.HandlerFunc(authHandler.LinkWallet)))
	mux.Handle("/api/v1/auth/characters", authMiddleware(http.HandlerFunc(authHandler.GetCharacters)))
	mux.Handle("/api/v1/auth/characters/create", authMiddleware(http.HandlerFunc(authHandler.CreateCharacter)))
	mux.Handle("/api/v1/vehicles/mint-starter", authMiddleware(http.HandlerFunc(vehicleHandler.MintStarter)))
	mux.Handle("/api/v1/vehicles", authMiddleware(http.HandlerFunc(vehicleHandler.ListVehicles)))
	mux.Handle("/api/v1/vehicles/cp", authMiddleware(http.HandlerFunc(vehicleHandler.GetVehicleCP)))
	mux.Handle("/api/v1/vehicles/equip", authMiddleware(http.HandlerFunc(vehicleHandler.EquipItem)))
	mux.Handle("/api/v1/vehicles/unequip", authMiddleware(http.HandlerFunc(vehicleHandler.UnequipItem)))
	
	// DDS & Items
	mux.Handle("/api/v1/items", authMiddleware(http.HandlerFunc(vehicleHandler.ListItems)))
	mux.Handle("/api/v1/items/damage", authMiddleware(http.HandlerFunc(vehicleHandler.ApplyDamage)))
	mux.Handle("/api/v1/items/repair", authMiddleware(http.HandlerFunc(vehicleHandler.Repair)))

	mux.Handle("/api/v1/combat/attack", authMiddleware(http.HandlerFunc(combatHandler.SimulateAttack)))
	mux.Handle("/api/v1/gacha/pull", authMiddleware(http.HandlerFunc(gachaHandler.Pull)))
	mux.Handle("/api/v1/exploration/start", authMiddleware(http.HandlerFunc(explorationHandler.StartExploration)))
	mux.Handle("/api/v1/exploration/timeline", authMiddleware(http.HandlerFunc(explorationHandler.GetTimeline)))
	mux.Handle("/api/v1/exploration/resolve", authMiddleware(http.HandlerFunc(explorationHandler.ResolveChoice)))
	mux.Handle("/api/v1/exploration/resolve-node", authMiddleware(http.HandlerFunc(explorationHandler.ResolveNode)))
	mux.Handle("/api/v1/exploration/advance", authMiddleware(http.HandlerFunc(explorationHandler.AdvanceTimeline)))
	mux.Handle("/api/v1/game/pilot-stats", authMiddleware(http.HandlerFunc(gameHandler.GetPilotStats)))

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
