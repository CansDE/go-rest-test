package handlers

import (
	"test-rest-api/internal/services"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	router  fiber.Router
	service *services.UserService
}

func NewUserHandler(router fiber.Router, service *services.UserService) *UserHandler {
	return &UserHandler{router: router, service: service}
}

func (u *UserHandler) RegisterRoutes() {
	u.router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(u.service.GetAll())
	})

	u.router.Post("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(u.service.Create(ctx))
	})
}
