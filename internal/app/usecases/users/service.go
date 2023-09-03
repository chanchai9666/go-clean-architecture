package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"eql/internal/app/entities/models"
	"eql/internal/app/entities/schema"
	"eql/internal/app/repositories"
)

type Service interface {
	GetUserAll(c *fiber.Ctx) ([]models.User, error)
	GetUser(c *fiber.Ctx, req *schema.UserRequest) ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewService(repo repositories.UserRepository) Service {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserAll(c *fiber.Ctx) ([]models.User, error) {
	user, err := s.repo.GetUserData()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUser(c *fiber.Ctx, req *schema.UserRequest) ([]models.User, error) {
	fmt.Println("asdfghj")
	user, err := s.repo.GetUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
