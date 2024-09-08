package router

import (
	"server/internal/api"
	"server/internal/repository"
	"server/internal/service"

	"github.com/gofiber/fiber/v2"
)

type routeStruct struct {
	app     *fiber.App
	handler *api.ServiceServer
}

func NewRouter(app *fiber.App, dao *repository.DAO) *routeStruct {
	exampleService := service.NewExampleService(dao)
	authService := service.NewAuthService(dao)
	perangkatService := service.NewPerangkatService(dao)
	objectService := service.NewObjectService(dao)
	handler := api.NewService(exampleService, authService, perangkatService, objectService)
	return &routeStruct{
		app:     app,
		handler: handler,
	}
}
