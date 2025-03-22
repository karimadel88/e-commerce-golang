package middleware

import (
	"context"
	"ecommerce-app/internal/service"
	"ecommerce-app/pkg/logger"
	"net/http"
	"strings"
)

// UserAuthKey is the key used to store the user ID in the request context
type UserAuthKey string

// UserIDKey is the key used to store the user ID in the request context
const UserIDKey UserAuthKey = "user_id"

// UserAuth middleware ensures that only authenticated users can access protected routes
func UserAuth(authService service.AuthService) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log := logger.New()

			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				log.Error("Missing Authorization header")
				return
			}

			// Check if the Authorization header has the correct format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				log.Error("Invalid Authorization header format")
				return
			}

			// Extract the token
			tokenString := parts[1]

			// Validate the token
			claims, err := authService.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				log.Error("Invalid token: " + err.Error())
				return
			}

			// Add the user ID to the request context
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			
			// Call the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

// GetUserID extracts the user ID from the request context
func GetUserID(r *http.Request) (uint, bool) {
	userID, ok := r.Context().Value(UserIDKey).(uint)
	return userID, ok
}