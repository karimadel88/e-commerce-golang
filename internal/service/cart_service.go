package service

import (
	"ecommerce-app/internal/models"
	"ecommerce-app/internal/repository"
)

// CartService defines the interface for cart-related business logic
type CartService interface {
	GetCart(userID uint) ([]models.CartItem, error)
	AddToCart(userID uint, productID uint, quantity int) error
	RemoveFromCart(userID uint, productID uint) error
}

// DefaultCartService implements CartService
type DefaultCartService struct {
	repo repository.CartRepository
}

// NewCartService creates a new instance of DefaultCartService
func NewCartService(repo repository.CartRepository) CartService {
	return &DefaultCartService{
		repo: repo,
	}
}

// GetCart retrieves a user's cart items
func (s *DefaultCartService) GetCart(userID uint) ([]models.CartItem, error) {
	return s.repo.GetCart(userID)
}

// AddToCart adds a product to the user's cart
func (s *DefaultCartService) AddToCart(userID uint, productID uint, quantity int) error {
	cartItem := &models.CartItem{
		CartID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}
	return s.repo.AddToCart(cartItem)
}

// RemoveFromCart removes a product from the user's cart
func (s *DefaultCartService) RemoveFromCart(userID uint, productID uint) error {
	return s.repo.RemoveFromCart(userID, productID)
}