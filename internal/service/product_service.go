package service

import (
    "ecommerce-app/internal/models"
    "ecommerce-app/internal/repository"
)

// ProductService defines the interface for product-related business logic
type ProductService interface {
    GetProductByID(id uint) (*models.Product, error)
    GetAllProducts() ([]models.Product, error)
    GetProductsPaginated(page, pageSize int) ([]models.Product, error)
    CreateProduct(product *models.Product) error
    UpdateProduct(product *models.Product) error
    DeleteProduct(id uint) error
    CountProducts() (int64, error)
}

// DefaultProductService implements ProductService
type DefaultProductService struct {
    repo repository.ProductRepository
}

// NewProductService creates a new instance of DefaultProductService
func NewProductService(repo repository.ProductRepository) ProductService {
    return &DefaultProductService{
        repo: repo,
    }
}

// GetProductByID retrieves a product by its ID
func (s *DefaultProductService) GetProductByID(id uint) (*models.Product, error) {
    return s.repo.FindByID(id)
}

// GetAllProducts retrieves all products
func (s *DefaultProductService) GetAllProducts() ([]models.Product, error) {
    return s.repo.List()
}

// GetProductsPaginated retrieves products with pagination
func (s *DefaultProductService) GetProductsPaginated(page, pageSize int) ([]models.Product, error) {
    return s.repo.ListPaginated(page, pageSize)
}

// CreateProduct creates a new product
func (s *DefaultProductService) CreateProduct(product *models.Product) error {
    return s.repo.Create(product)
}

// UpdateProduct updates an existing product
func (s *DefaultProductService) UpdateProduct(product *models.Product) error {
    return s.repo.Update(product)
}

// DeleteProduct deletes a product by its ID
func (s *DefaultProductService) DeleteProduct(id uint) error {
    return s.repo.Delete(id)
}

// CountProducts returns the total number of products
func (s *DefaultProductService) CountProducts() (int64, error) {
    return s.repo.Count()
}