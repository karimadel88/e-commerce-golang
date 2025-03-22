package handlers

import (
	"ecommerce-app/internal/service"
	"ecommerce-app/pkg/logger"
	"encoding/json"
	"net/http"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authService service.AuthService
	log         *logger.Logger
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         logger.New(),
	}
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ResetPasswordRequestRequest represents the request body for password reset request
type ResetPasswordRequestRequest struct {
	Email string `json:"email"`
}

// ResetPasswordRequest represents the request body for password reset
type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Error("Failed to parse request body: " + err.Error())
		responseWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Email == "" || req.Password == "" {
		responseWithError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Register user
	user, err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		h.log.Error("Failed to register user: " + err.Error())
		responseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success response
	responseWithJSON(w, AuthResponse{
		Success: true,
		Message: "User registered successfully",
		Data: map[string]interface{}{
			"user_id": user.ID,
			"email":   user.Email,
		},
	}, http.StatusCreated)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Error("Failed to parse request body: " + err.Error())
		responseWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Email == "" || req.Password == "" {
		responseWithError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Login user
	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.log.Error("Failed to login user: " + err.Error())
		responseWithError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Return success response with token
	responseWithJSON(w, AuthResponse{
		Success: true,
		Message: "Login successful",
		Data: map[string]string{
			"token": token,
		},
	}, http.StatusOK)
}

// RequestPasswordReset handles password reset requests
func (h *AuthHandler) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req ResetPasswordRequestRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Error("Failed to parse request body: " + err.Error())
		responseWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Email == "" {
		responseWithError(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Request password reset
	err = h.authService.ResetPasswordRequest(req.Email)
	if err != nil {
		h.log.Error("Failed to request password reset: " + err.Error())
		responseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Always return success to prevent email enumeration
	responseWithJSON(w, AuthResponse{
		Success: true,
		Message: "If your email is registered, you will receive a password reset link",
	}, http.StatusOK)
}

// ResetPassword handles password reset
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req ResetPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Error("Failed to parse request body: " + err.Error())
		responseWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Token == "" || req.NewPassword == "" {
		responseWithError(w, "Token and new password are required", http.StatusBadRequest)
		return
	}

	// Reset password
	err = h.authService.ResetPassword(req.Token, req.NewPassword)
	if err != nil {
		h.log.Error("Failed to reset password: " + err.Error())
		responseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success response
	responseWithJSON(w, AuthResponse{
		Success: true,
		Message: "Password reset successful",
	}, http.StatusOK)
}

// Helper function to send JSON response
func responseWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// Helper function to send error response
func responseWithError(w http.ResponseWriter, message string, statusCode int) {
	responseWithJSON(w, AuthResponse{
		Success: false,
		Message: message,
	}, statusCode)
}