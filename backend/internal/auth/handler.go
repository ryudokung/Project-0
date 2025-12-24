package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Login request received: %s %s", r.Method, r.URL.Path)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Also check Authorization header for Privy Token
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		req.PrivyToken = authHeader[7:]
	}

	resp, err := h.useCase.Login(req)
	if err != nil {
		log.Printf("Login error: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[7:]
	user, err := h.useCase.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.useCase.Signup(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) LinkWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserID        string `json:"user_id"`
		WalletAddress string `json:"wallet_address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.useCase.LinkWallet(uid, req.WalletAddress); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *Handler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user from token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[7:]
	user, err := h.useCase.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req CreateCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	char, err := h.useCase.CreateCharacter(user.ID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(char)
}

func (h *Handler) GetCharacters(w http.ResponseWriter, r *http.Request) {
	// Get user from token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[7:]
	user, err := h.useCase.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	chars, err := h.useCase.GetCharacters(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chars)
}
