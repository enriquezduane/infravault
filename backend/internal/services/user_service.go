package services

import "infravault/internal/repositories"

type UserService struct {
	userRepo repositories.UserRepository
}
