package repository

import (
    "ecommerce-app/internal/models"
    "gorm.io/gorm"
)

// OrderRepository defines the interface for order-related database operations
type OrderRepository interface {
    Create(order *models.Order) error
    FindByID(id uint) (*models.Order, error)
    Update(order *models.Order) error
    UpdateStatus(id uint, status string) error
    Delete(id uint) error
    List() ([]models.Order, error)
    ListPaginated(page, pageSize int) ([]models.Order, error)
    ListWithUser() ([]models.Order, error)
    ListWithUserPaginated(page, pageSize int) ([]models.Order, error)
    Count() (int64, error)
}

// GormOrderRepository implements OrderRepository using GORM
type GormOrderRepository struct {
    db *gorm.DB
}

// NewOrderRepository creates a new instance of GormOrderRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
    return &GormOrderRepository{
        db: db,
    }
}

// Create inserts a new order into the database
func (r *GormOrderRepository) Create(order *models.Order) error {
    return r.db.Create(order).Error
}

// FindByID retrieves an order by its ID
func (r *GormOrderRepository) FindByID(id uint) (*models.Order, error) {
    var order models.Order
    err := r.db.First(&order, id).Error
    if err != nil {
        return nil, err
    }
    return &order, nil
}

// Update modifies an existing order in the database
func (r *GormOrderRepository) Update(order *models.Order) error {
    return r.db.Save(order).Error
}

// UpdateStatus updates just the status field of an order
func (r *GormOrderRepository) UpdateStatus(id uint, status string) error {
    return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

// Delete removes an order from the database
func (r *GormOrderRepository) Delete(id uint) error {
    return r.db.Delete(&models.Order{}, id).Error
}

// List retrieves all orders
func (r *GormOrderRepository) List() ([]models.Order, error) {
    var orders []models.Order
    err := r.db.Find(&orders).Error
    return orders, err
}

// ListPaginated retrieves orders with pagination
func (r *GormOrderRepository) ListPaginated(page, pageSize int) ([]models.Order, error) {
    var orders []models.Order
    offset := (page - 1) * pageSize
    err := r.db.Offset(offset).Limit(pageSize).Find(&orders).Error
    return orders, err
}

// ListWithUser retrieves all orders with preloaded user data
func (r *GormOrderRepository) ListWithUser() ([]models.Order, error) {
    var orders []models.Order
    err := r.db.Preload("User").Find(&orders).Error
    return orders, err
}

// ListWithUserPaginated retrieves orders with preloaded user data and pagination
func (r *GormOrderRepository) ListWithUserPaginated(page, pageSize int) ([]models.Order, error) {
    var orders []models.Order
    offset := (page - 1) * pageSize
    err := r.db.Preload("User").Offset(offset).Limit(pageSize).Find(&orders).Error
    return orders, err
}

// Count returns the total number of orders
func (r *GormOrderRepository) Count() (int64, error) {
    var count int64
    err := r.db.Model(&models.Order{}).Count(&count).Error
    return count, err
}