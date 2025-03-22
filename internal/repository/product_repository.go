package repository

import (
	"ecommerce-app/internal/models"

	"gorm.io/gorm"
)

// ProductRepository defines the interface for product-related database operations
type ProductRepository interface {
    Create(product *models.Product) error
    FindByID(id uint) (*models.Product, error)
    FindByCategory(category string) ([]models.Product, error)
    SearchByName(name string) ([]models.Product, error)
    Update(product *models.Product) error
    Delete(id uint) error
    List() ([]models.Product, error)
    ListPaginated(page, pageSize int) ([]models.Product, error)
    Count() (int64, error)
}

// GormProductRepository implements ProductRepository using GORM
type GormProductRepository struct {
    db *gorm.DB
}

// NewProductRepository creates a new instance of GormProductRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
    return &GormProductRepository{
        db: db,
    }
}

// Create inserts a new product into the database
func (r *GormProductRepository) Create(product *models.Product) error {
    return r.db.Create(product).Error
}

// FindByID retrieves a product by its ID
func (r *GormProductRepository) FindByID(id uint) (*models.Product, error) {
    var product models.Product
    err := r.db.First(&product, id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

// Update modifies an existing product in the database
func (r *GormProductRepository) Update(product *models.Product) error {
    return r.db.Save(product).Error
}

// Delete removes a product from the database
func (r *GormProductRepository) Delete(id uint) error {
    return r.db.Delete(&models.Product{}, id).Error
}

// List retrieves all products
func (r *GormProductRepository) List() ([]models.Product, error) {
    var products []models.Product
    err := r.db.Find(&products).Error
    return products, err
}

// ListPaginated retrieves products with pagination
func (r *GormProductRepository) ListPaginated(page, pageSize int) ([]models.Product, error) {
    var products []models.Product
    offset := (page - 1) * pageSize
    err := r.db.Offset(offset).Limit(pageSize).Find(&products).Error
    return products, err
}

// FindByCategory retrieves products by category
func (r *GormProductRepository) FindByCategory(category string) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Where("category = ?", category).Find(&products).Error
    return products, err
}

// SearchByName searches for products with names containing the search term
func (r *GormProductRepository) SearchByName(name string) ([]models.Product, error) {
    var products []models.Product
    searchTerm := "%" + name + "%"
    err := r.db.Where("name ILIKE ?", searchTerm).Find(&products).Error
    return products, err
}

// Count returns the total number of products with optimized query
func (r *GormProductRepository) Count() (int64, error) {
    var count int64
    // Using a more efficient counting query that doesn't load the entire table
    err := r.db.Table("products").Count(&count).Error
    return count, err
}