package services

import (
	"test-rest-api/internal/models"
	"test-rest-api/internal/repository"

	"github.com/gofiber/fiber/v3"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) GetAll() []models.User {
	return u.repository.GetAll()
}

func (u *UserService) Create(ctx fiber.Ctx) *models.User {
	user := models.UserFromContext(ctx)
	u.repository.Create(user)
	return user
}
