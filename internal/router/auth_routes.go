package router

import (
	"ecommerce-app/internal/handlers"
	"net/http"
)

// SetupAuthRoutes configures authentication-related routes
func SetupAuthRoutes(authHandler *handlers.AuthHandler) {
	// Authentication routes
	http.HandleFunc("/auth/register", authHandler.Register)
	http.HandleFunc("/auth/login", authHandler.Login)
	http.HandleFunc("/auth/reset-password-request", authHandler.RequestPasswordReset)
	http.HandleFunc("/auth/reset-password", authHandler.ResetPassword)
}