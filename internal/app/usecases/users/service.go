package users

import (
	"github.com/gofiber/fiber/v2"

	"eql/internal/app/entities/models"
	"eql/internal/app/repositories"
)

type Service interface {
	GetUser(c *fiber.Ctx) ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewService(repo repositories.UserRepository) Service {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUser(c *fiber.Ctx) ([]models.User, error) {
	user, err := s.repo.GetUserData()
	if err != nil {
		return nil, err
	}
	return user, nil
}
