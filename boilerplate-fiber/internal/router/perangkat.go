package router

import (
	"server/internal/middleware"
)

func (m *routeStruct) Perangkat() {
	subRoute := m.app.Group("/perangkat")
	subRoute.Use(middleware.Jwt())
	subRoute.Use(middleware.Permission())
	subRoute.Get("/", m.handler.PerangkatList)
	subRoute.Post("/", m.handler.PerangkatCreate)
	subRoute.Get("/:id/", m.handler.PerangkatDetail)
	subRoute.Patch("/:id/", m.handler.PerangkatPatch)
	subRoute.Patch("/:id/reset-password/", m.handler.PerangkatResetPassword)
	subRoute.Delete("/:id/", m.handler.PerangkatDelete)
}
