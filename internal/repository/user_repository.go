package repository

import (
    "ecommerce-app/internal/models"
    "gorm.io/gorm"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
    Create(user *models.User) error
    FindByID(id uint) (*models.User, error)
    FindByEmail(email string) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
    List() ([]models.User, error)
    ListPaginated(page, pageSize int) ([]models.User, error)
    Count() (int64, error)
}

// GormUserRepository implements UserRepository using GORM
type GormUserRepository struct {
    db *gorm.DB
}

// NewUserRepository creates a new instance of GormUserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
    return &GormUserRepository{
        db: db,
    }
}

// Create inserts a new user into the database
func (r *GormUserRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

// FindByID retrieves a user by its ID
func (r *GormUserRepository) FindByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// FindByEmail retrieves a user by email
func (r *GormUserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// Update modifies an existing user in the database
func (r *GormUserRepository) Update(user *models.User) error {
    return r.db.Save(user).Error
}

// Delete removes a user from the database
func (r *GormUserRepository) Delete(id uint) error {
    return r.db.Delete(&models.User{}, id).Error
}

// List retrieves all users
func (r *GormUserRepository) List() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error
    return users, err
}

// ListPaginated retrieves users with pagination
func (r *GormUserRepository) ListPaginated(page, pageSize int) ([]models.User, error) {
    var users []models.User
    offset := (page - 1) * pageSize
    err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error
    return users, err
}

// Count returns the total number of users
func (r *GormUserRepository) Count() (int64, error) {
    var count int64
    err := r.db.Model(&models.User{}).Count(&count).Error
    return count, err
}