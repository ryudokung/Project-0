package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/ryudokung/Project-0/backend/internal/auth/constants"
)

func Middleware(useCase UseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := authHeader[7:]
			user, err := useCase.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			if user == nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}

			ctx := context.WithValue(r.Context(), constants.UserIDKey, user.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
