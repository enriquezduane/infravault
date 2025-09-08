package repositories

import "infravault/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByGoogleID(googleID string) (*models.User, error)
	GetByID(id int) (*models.User, error)
}
