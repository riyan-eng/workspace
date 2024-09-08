package main

import (
	"fmt"
	"runtime"
	"server/config"
	"server/env"
	"server/infrastructure"
	"server/internal/repository"
	"server/internal/router"

	_ "server/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger" // hertz-swagger middleware
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU - 1)
	}
	env.LoadEnvironmentFile()
	env.NewEnv()

	config.NewLimiterStore()
	config.NewLogger()

	infrastructure.ConnectSqlDB()
	infrastructure.ConnectSqlxDB()
	infrastructure.ConnRedis()
	infrastructure.NewLocalizer()
}

// @title Sisambi
// @version 1.0
// @description This is a Sisambi Api Documentation.

// @contact.name hertz-contrib
// @contact.url https://github.com/hertz-contrib

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer access token here
func main() {
	// create instance
	app := fiber.New()

	// swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	// middleware
	app.Use(recover.New())
	// app.Use(middleware.RequestId())
	// app.Use(middleware.Limiter())
	app.Use(infrastructure.LocalizerMiddleware())

	// corsconfig := cors.DefaultConfig()
	// corsconfig.AllowCredentials = true
	// corsconfig.AllowOrigins = []string{"*"}
	// corsconfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	// app.Use(cors.New(corsconfig))

	// service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.SqlxDB, infrastructure.Redis, config.NewEnforcer())

	// router
	routers := router.NewRouter(app, &dao)
	routers.Index()
	routers.Authentication()
	routers.Example()
	routers.Perangkat()
	routers.Object()

	// startup log
	fmt.Println("server run on:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)

	app.Listen(":3000")
}
