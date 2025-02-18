package infrastructure

import (
	"test/internal/application"

	"github.com/gofiber/fiber/v2"
)

func InitializeRouter(app *application.Application) *fiber.App {
	dep := application.SetupDependencyInjection(app)

	v1 := app.App.Group("api")
	{

	}

	return app.App
}
