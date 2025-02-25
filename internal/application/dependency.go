package application

import (
	"test/internal/controller"
	"test/internal/repository"
	"test/internal/service"
)

type Dependency struct {
	AuthController   controller.AuthController
	AnimalController controller.AnimalController
}

func SetupDependencyInjection(app *Application) *Dependency {
	//repo
	authRepo := repository.NewAuthRepository(app.Context, app.Config, app.DB)
	animalRepo := repository.NewAnimalRepository(app.Context, app.Config, app.DB)
	//svc
	authSvc := service.NewAuthService(app.Context, app.Config, authRepo)
	animalSvc := service.NewAnimalService(app.Context, app.Config, animalRepo)
	//ctrl
	authCtrl := controller.NewAuthController(app.Context, app.Config, authSvc)
	animalCtrl := controller.NewAnimalController(app.Context, app.Config, animalSvc)

	return &Dependency{
		AuthController:   authCtrl,
		AnimalController: animalCtrl,
	}
}
