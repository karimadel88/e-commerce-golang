package service

import (
    "ecommerce-app/internal/models"
    "ecommerce-app/internal/repository"
)

// UserService defines the interface for user-related business logic
type UserService interface {
    GetUserByID(id uint) (*models.User, error)
    GetUserByEmail(email string) (*models.User, error)
    GetAllUsers() ([]models.User, error)
    GetUsersPaginated(page, pageSize int) ([]models.User, error)
    CreateUser(user *models.User) error
    UpdateUser(user *models.User) error
    DeleteUser(id uint) error
    CountUsers() (int64, error)
}

// DefaultUserService implements UserService
type DefaultUserService struct {
    repo repository.UserRepository
}

// NewUserService creates a new instance of DefaultUserService
func NewUserService(repo repository.UserRepository) UserService {
    return &DefaultUserService{
        repo: repo,
    }
}

// GetUserByID retrieves a user by its ID
func (s *DefaultUserService) GetUserByID(id uint) (*models.User, error) {
    return s.repo.FindByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *DefaultUserService) GetUserByEmail(email string) (*models.User, error) {
    return s.repo.FindByEmail(email)
}

// GetAllUsers retrieves all users
func (s *DefaultUserService) GetAllUsers() ([]models.User, error) {
    return s.repo.List()
}

// GetUsersPaginated retrieves users with pagination
func (s *DefaultUserService) GetUsersPaginated(page, pageSize int) ([]models.User, error) {
    return s.repo.ListPaginated(page, pageSize)
}

// CreateUser creates a new user
func (s *DefaultUserService) CreateUser(user *models.User) error {
    return s.repo.Create(user)
}

// UpdateUser updates an existing user
func (s *DefaultUserService) UpdateUser(user *models.User) error {
    return s.repo.Update(user)
}

// DeleteUser deletes a user by its ID
func (s *DefaultUserService) DeleteUser(id uint) error {
    return s.repo.Delete(id)
}

// CountUsers returns the total number of users
func (s *DefaultUserService) CountUsers() (int64, error) {
    return s.repo.Count()
}