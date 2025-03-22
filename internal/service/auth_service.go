package service

import (
	"ecommerce-app/internal/models"
	"ecommerce-app/pkg/logger"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines the interface for authentication-related business logic
type AuthService interface {
	Register(email, password string) (*models.User, error)
	Login(email, password string) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
	ResetPasswordRequest(email string) error
	ResetPassword(token, newPassword string) error
	SeedTestUser(email, password string) (*models.User, error)
}

// TokenClaims represents the JWT claims
type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// DefaultAuthService implements AuthService
type DefaultAuthService struct {
	userService UserService
	jwtSecret   []byte
	log         *logger.Logger
}

// NewAuthService creates a new instance of DefaultAuthService
func NewAuthService(userService UserService) AuthService {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret_key" // Not recommended for production
	}

	return &DefaultAuthService{
		userService: userService,
		jwtSecret:   []byte(jwtSecret),
		log:         logger.New(),
	}
}

// Register creates a new user account
func (s *DefaultAuthService) Register(email, password string) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.userService.GetUserByEmail(email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Failed to hash password: " + err.Error())
		return nil, errors.New("failed to create user")
	}

	// Create new user
	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userService.CreateUser(user)
	if err != nil {
		s.log.Error("Failed to create user: " + err.Error())
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

// Login authenticates a user and returns a JWT token
func (s *DefaultAuthService) Login(email, password string) (string, error) {
	// Find user by email
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	expTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := &TokenClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		s.log.Error("Failed to generate token: " + err.Error())
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *DefaultAuthService) ValidateToken(tokenString string) (*TokenClaims, error) {
	claims := &TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ResetPasswordRequest initiates a password reset process
func (s *DefaultAuthService) ResetPasswordRequest(email string) error {
	// Find user by email
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		// Don't reveal that the email doesn't exist
		return nil
	}

	// Generate a random token
	token := generateRandomToken(32)

	// Set token expiry (24 hours from now)
	expiry := time.Now().Add(24 * time.Hour)

	// Update user with reset token and expiry
	user.ResetToken = &token
	user.ResetTokenExpiry = &expiry

	err = s.userService.UpdateUser(user)
	if err != nil {
		s.log.Error("Failed to update user with reset token: " + err.Error())
		return errors.New("failed to process password reset")
	}

	// In a real application, you would send an email with the reset link
	// For this example, we'll just log it
	s.log.Info("Password reset token for " + email + ": " + token)

	return nil
}

// ResetPassword completes the password reset process
func (s *DefaultAuthService) ResetPassword(token, newPassword string) error {
	// Find all users
	users, err := s.userService.GetAllUsers()
	if err != nil {
		s.log.Error("Failed to get users: " + err.Error())
		return errors.New("failed to process password reset")
	}

	// Find user with matching reset token
	var user *models.User
	for i := range users {
		if users[i].ResetToken != nil && *users[i].ResetToken == token {
			user = &users[i]
			break
		}
	}

	if user == nil {
		return errors.New("invalid or expired reset token")
	}

	// Check if token is expired
	if user.ResetTokenExpiry == nil || user.ResetTokenExpiry.Before(time.Now()) {
		return errors.New("invalid or expired reset token")
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Failed to hash password: " + err.Error())
		return errors.New("failed to reset password")
	}

	// Update user with new password and clear reset token
	user.PasswordHash = string(hashedPassword)
	user.ResetToken = nil
	user.ResetTokenExpiry = nil

	err = s.userService.UpdateUser(user)
	if err != nil {
		s.log.Error("Failed to update user with new password: " + err.Error())
		return errors.New("failed to reset password")
	}

	return nil
}

// Helper function to generate a random token
func generateRandomToken(length int) string {
	// In a real application, use a secure random generator
	// For this example, we'll use a simple timestamp-based token
	return time.Now().Format("20060102150405") + "token"
}

// SeedTestUser creates a test user with the specified email and password if it doesn't exist
func (s *DefaultAuthService) SeedTestUser(email, password string) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.userService.GetUserByEmail(email)
	if err == nil && existingUser != nil {
		// User already exists, return it
		s.log.Info("Test user already exists: " + email)
		return existingUser, nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Failed to hash password for test user: " + err.Error())
		return nil, errors.New("failed to create test user")
	}

	// Create new test user
	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userService.CreateUser(user)
	if err != nil {
		s.log.Error("Failed to create test user: " + err.Error())
		return nil, errors.New("failed to create test user")
	}

	s.log.Info("Successfully created test user: " + email)
	return user, nil
}