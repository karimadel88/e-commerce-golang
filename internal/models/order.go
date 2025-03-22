package models

import (
    "time"
    "gorm.io/gorm"
)

// Order represents the order model in the database
type Order struct {
    ID         uint           `gorm:"primaryKey"`
    UserID     uint           `gorm:"not null"`
    User       User           `gorm:"foreignKey:UserID"`
    Total      float64        `gorm:"type:decimal(10,2);not null"`
    Status     string         `gorm:"type:varchar(50);default:pending"`
    OrderItems []OrderItem    `gorm:"foreignKey:OrderID"`
    CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
}

// BeforeUpdate will be called before updating the order
func (o *Order) BeforeUpdate(tx *gorm.DB) error {
    o.UpdatedAt = time.Now()
    return nil
}

// OrderItem represents the order item model in the database
type OrderItem struct {
    ID           uint           `gorm:"primaryKey"`
    OrderID      uint           `gorm:"not null"`
    Order        Order          `gorm:"foreignKey:OrderID"`
    ProductID    uint           `gorm:"not null"`
    Product      Product        `gorm:"foreignKey:ProductID"`
    Quantity     int            `gorm:"not null"`
    PriceAtTime  float64        `gorm:"type:decimal(10,2);not null"`
    CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
}