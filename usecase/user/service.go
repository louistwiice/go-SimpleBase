package user

import (
	"github.com/louistwiice/go/simplebase/domain"
	"github.com/louistwiice/go/simplebase/models"
)

type service struct {
	repo domain.UserRepository
}

func NewUSerService(r domain.UserRepository) *service {
	return &service{
		repo: r,
	}
}

// List all users
func (s *service) List() ([]*models.User, error) {
	return s.repo.List()
}

// Create a user
func (s *service) Create(u *models.User) error {
	err := u.SetPassword()
	if err != nil {
		return err
	}
	return s.repo.Create(u)
}

// Retrieve a user
func (s *service) Get(id string) (*models.User, error) {
	u, err := s.repo.Get(id)
	if err != nil {
		return &models.User{}, models.ErrNotFound
	}
	return u, nil
}

// Update a user
func (s *service) Update(u *models.User) error {
	err := u.Validate()
	if err != nil {
		return err
	}
	return s.repo.Update(u)
}

// Update user password
func (s *service) UpdatePassword(u *models.User) error {
	err := u.SetPassword()
	if err != nil {
		return err
	}
	return s.repo.UpdatePassword(u)
}

// Delete a user
func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
