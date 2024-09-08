package router

import (
	"server/internal/api"
	"server/internal/repository"
	"server/internal/service"
	rpcserver "server/rpc_server"

	"github.com/gin-gonic/gin"
)

type routeStruct struct {
	app     *gin.Engine
	handler *api.ServiceServer
}

func NewRouter(app *gin.Engine, dao *repository.DAO) *routeStruct {
	exampleService := service.NewExampleService(dao)
	authService := service.NewAuthService(dao)
	perangkatService := service.NewPerangkatService(dao)
	objectService := service.NewObjectService(dao)

	exampleRpcServer := rpcserver.ExampleService()
	authRpcServer := rpcserver.AuthService()

	handler := api.NewService(
		exampleService,
		authService,
		perangkatService,
		objectService,
		exampleRpcServer,
		authRpcServer,
	)
	return &routeStruct{
		app:     app,
		handler: handler,
	}
}
