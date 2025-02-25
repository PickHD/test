package infrastructure

import (
	"test/internal/application"
	"test/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitializeRouter(app *application.Application) *fiber.App {
	dep := application.SetupDependencyInjection(app)

	v1 := app.App.Group("api")
	{
		v1.Post("/auth/register", dep.AuthController.Register)
		v1.Post("/auth/login", dep.AuthController.Login)

		v1.Post("/animal", middleware.ValidateJWTMiddleware, dep.AnimalController.Create)
		v1.Get("/animal", middleware.ValidateJWTMiddleware, dep.AnimalController.GetAll)
	}

	return app.App
}
