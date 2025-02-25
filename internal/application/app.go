package application

import (
	"context"
	"fmt"
	"log"
	"test/internal/config"
	"test/internal/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	App     *fiber.App
	Context context.Context
	Config  *config.Configuration
	DB      *gorm.DB
}

func NewApplication() (*Application, error) {
	app := &Application{}
	app.Context = context.TODO()
	app.Config = config.NewConfig()

	// database
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", app.Config.Database.Host, app.Config.Database.Username, app.Config.Database.Password, app.Config.Database.DbName, app.Config.Database.Port)), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return app, err
	}
	app.DB = db

	app.DB.AutoMigrate(
		&model.Animal{},
		&model.User{},
	)

	app.App = fiber.New()

	return app, nil
}
