package repository

import (
    "gorm.io/gorm"
)

// Repository provides a generic interface for database operations
type Repository[T any] interface {
    Create(entity *T) error
    FindByID(id uint) (*T, error)
    Update(entity *T) error
    Delete(id uint) error
    List(offset, limit int) ([]T, error)
}

// BaseRepository implements the Repository interface
type BaseRepository[T any] struct {
    db *gorm.DB
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
    return &BaseRepository[T]{
        db: db,
    }
}

// Create inserts a new entity into the database
func (r *BaseRepository[T]) Create(entity *T) error {
    return r.db.Create(entity).Error
}

// FindByID retrieves an entity by its ID
func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
    var entity T
    err := r.db.First(&entity, id).Error
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

// Update modifies an existing entity in the database
func (r *BaseRepository[T]) Update(entity *T) error {
    return r.db.Save(entity).Error
}

// Delete removes an entity from the database
func (r *BaseRepository[T]) Delete(id uint) error {
    var entity T
    return r.db.Delete(&entity, id).Error
}

// List retrieves a paginated list of entities
func (r *BaseRepository[T]) List(offset, limit int) ([]T, error) {
    var entities []T
    err := r.db.Offset(offset).Limit(limit).Find(&entities).Error
    return entities, err
}