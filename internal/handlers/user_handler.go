package handlers

import (
	"test-rest-api/internal/models"
	"test-rest-api/internal/repository"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	router     fiber.Router
	repository *repository.UserRepository
}

func NewUserHandler(router fiber.Router, repository *repository.UserRepository) *UserHandler {
	return &UserHandler{router: router, repository: repository}
}

func (u *UserHandler) RegisterRoutes() {
	u.router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(u.repository.GetAll())
	})

	u.router.Post("/", func(ctx fiber.Ctx) error {
		user := models.UserFromContext(ctx)
		u.repository.Create(user)
		return ctx.JSON(user)
	})
}
