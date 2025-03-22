package models

import (
    "time"
    "gorm.io/gorm"
)

// User represents the user model in the database
type User struct {
    ID               uint           `gorm:"primaryKey"`
    Email            string         `gorm:"type:varchar(255);unique;not null"`
    PasswordHash     string         `gorm:"type:varchar(255);not null"`
    ResetToken       *string        `gorm:"type:varchar(255)"`
    ResetTokenExpiry *time.Time     `gorm:"type:timestamp"`
    CreatedAt        time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt        time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    Carts            []Cart         `gorm:"foreignKey:UserID"`
    Orders           []Order        `gorm:"foreignKey:UserID"`
}

// BeforeUpdate will be called before updating the user
func (u *User) BeforeUpdate(tx *gorm.DB) error {
    u.UpdatedAt = time.Now()
    return nil
}