package server

import (
	"test-rest-api/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func StartServer() error {
	app := fiber.New()

	//ROUTES
	usersRoutes := app.Group("/users")

	//UserHandler irgendwie hier raus machen damit ich nicht gorm.DB in StartServer packen muss
	//für routes

	userHandler := handlers.NewUserHandler(usersRoutes)
	userHandler.RegisterRoutes()

	err := app.Listen(":8080")
	if err != nil {
		return err
	}
	return nil
}
