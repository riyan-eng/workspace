package router

import (
	"server/internal/middleware"
)

func (m *routeStruct) Perangkat() {
	subRoute := m.app.Group("/perangkat")
	subRoute.Use(middleware.Jwt())
	subRoute.Use(middleware.Permission())
	subRoute.GET("/", m.handler.PerangkatList)
	subRoute.POST("/", m.handler.PerangkatCreate)
	subRoute.GET("/:id/", m.handler.PerangkatDetail)
	subRoute.PATCH("/:id/", m.handler.PerangkatPatch)
	subRoute.PATCH("/:id/reset-password/", m.handler.PerangkatResetPassword)
	subRoute.DELETE("/:id/", m.handler.PerangkatDelete)
}
