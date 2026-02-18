package server

import (
	"test-rest-api/internal/handlers"
	"test-rest-api/internal/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) error {
	app := fiber.New()

	//ROUTES
	usersRoutes := app.Group("/users")

	//REPOS
	userRepo := repository.NewUserRepository(db)

	//HANDLERS
	userHandler := handlers.NewUserHandler(usersRoutes, userRepo)
	userHandler.RegisterRoutes()

	err := app.Listen(":8080")
	if err != nil {
		return err
	}
	return nil
}
