package handlers

import "github.com/gofiber/fiber/v3"

type UserHandler struct {
	router fiber.Router
}

func NewUserHandler(router fiber.Router) *UserHandler {
	return &UserHandler{router: router}
}

func (u *UserHandler) RegisterRoutes() {
	u.router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON(23)
	})
}
