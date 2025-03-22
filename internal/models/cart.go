package models

import (
    "time"
    "gorm.io/gorm"
)

// Cart represents the cart model in the database
type Cart struct {
    ID        uint           `gorm:"primaryKey"`
    UserID    uint           `gorm:"not null"`
    User      User           `gorm:"foreignKey:UserID"`
    CartItems []CartItem     `gorm:"foreignKey:CartID"`
    CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
}

// BeforeUpdate will be called before updating the cart
func (c *Cart) BeforeUpdate(tx *gorm.DB) error {
    c.UpdatedAt = time.Now()
    return nil
}

// CartItem represents the cart item model in the database
type CartItem struct {
    ID        uint           `gorm:"primaryKey"`
    CartID    uint           `gorm:"not null"`
    Cart      Cart           `gorm:"foreignKey:CartID"`
    ProductID uint           `gorm:"not null"`
    Product   Product        `gorm:"foreignKey:ProductID"`
    Quantity  int            `gorm:"not null;check:quantity > 0"`
    CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
}

// BeforeUpdate will be called before updating the cart item
func (ci *CartItem) BeforeUpdate(tx *gorm.DB) error {
    ci.UpdatedAt = time.Now()
    return nil
}