package models

import (
    "time"
    "gorm.io/gorm"
)

// Product represents the product model in the database
type Product struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string         `gorm:"type:varchar(255);not null"`
    Description string         `gorm:"type:text"`
    Price       float64        `gorm:"type:decimal(10,2);not null"`
    Stock       int            `gorm:"not null"`
    CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    CartItems   []CartItem     `gorm:"foreignKey:ProductID"`
    OrderItems  []OrderItem    `gorm:"foreignKey:ProductID"`
}

// BeforeUpdate will be called before updating the product
func (p *Product) BeforeUpdate(tx *gorm.DB) error {
    p.UpdatedAt = time.Now()
    return nil
}