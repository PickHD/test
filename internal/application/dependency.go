package application

type Dependency struct{}

func SetupDependencyInjection(app *Application) *Dependency { return &Dependency{} }
