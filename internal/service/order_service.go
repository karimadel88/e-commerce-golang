package service

import (
    "ecommerce-app/internal/models"
    "ecommerce-app/internal/repository"
)

// OrderService defines the interface for order-related business logic
type OrderService interface {
    GetOrderByID(id uint) (*models.Order, error)
    GetAllOrders() ([]models.Order, error)
    GetOrdersPaginated(page, pageSize int) ([]models.Order, error)
    GetAllOrdersWithUser() ([]models.Order, error)
    GetOrdersWithUserPaginated(page, pageSize int) ([]models.Order, error)
    CreateOrder(order *models.Order) error
    UpdateOrder(order *models.Order) error
    UpdateOrderStatus(id uint, status string) error
    DeleteOrder(id uint) error
    CountOrders() (int64, error)
}

// DefaultOrderService implements OrderService
type DefaultOrderService struct {
    repo repository.OrderRepository
}

// NewOrderService creates a new instance of DefaultOrderService
func NewOrderService(repo repository.OrderRepository) OrderService {
    return &DefaultOrderService{
        repo: repo,
    }
}

// GetOrderByID retrieves an order by its ID
func (s *DefaultOrderService) GetOrderByID(id uint) (*models.Order, error) {
    return s.repo.FindByID(id)
}

// GetAllOrders retrieves all orders
func (s *DefaultOrderService) GetAllOrders() ([]models.Order, error) {
    return s.repo.List()
}

// GetOrdersPaginated retrieves orders with pagination
func (s *DefaultOrderService) GetOrdersPaginated(page, pageSize int) ([]models.Order, error) {
    return s.repo.ListPaginated(page, pageSize)
}

// GetAllOrdersWithUser retrieves all orders with user information
func (s *DefaultOrderService) GetAllOrdersWithUser() ([]models.Order, error) {
    return s.repo.ListWithUser()
}

// GetOrdersWithUserPaginated retrieves orders with user information and pagination
func (s *DefaultOrderService) GetOrdersWithUserPaginated(page, pageSize int) ([]models.Order, error) {
    return s.repo.ListWithUserPaginated(page, pageSize)
}

// CreateOrder creates a new order
func (s *DefaultOrderService) CreateOrder(order *models.Order) error {
    return s.repo.Create(order)
}

// UpdateOrder updates an existing order
func (s *DefaultOrderService) UpdateOrder(order *models.Order) error {
    return s.repo.Update(order)
}

// UpdateOrderStatus updates just the status of an order
func (s *DefaultOrderService) UpdateOrderStatus(id uint, status string) error {
    return s.repo.UpdateStatus(id, status)
}

// DeleteOrder deletes an order by its ID
func (s *DefaultOrderService) DeleteOrder(id uint) error {
    return s.repo.Delete(id)
}

// CountOrders returns the total number of orders
func (s *DefaultOrderService) CountOrders() (int64, error) {
    return s.repo.Count()
}