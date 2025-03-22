package repository

import (
	"ecommerce-app/internal/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCart(userID uint) ([]models.CartItem, error)
	AddToCart(cart *models.CartItem) error
	RemoveFromCart(userID uint, productID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetCart(userID uint) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error
	return cartItems, err
}

func (r *cartRepository) AddToCart(cart *models.CartItem) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) RemoveFromCart(userID uint, productID uint) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.CartItem{}).Error
}