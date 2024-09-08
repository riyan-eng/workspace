package router

import (
	"server/internal/middleware"
)

func (m *routeStruct) Authentication() {
	subRoute := m.app.Group("/auth")
	subRoute.Post("/login/", m.handler.AuthLogin)
	subRoute.Post("/refresh/", m.handler.AuthRefresh)
	subRoute.Use(middleware.Jwt())
	subRoute.Get("/me/", m.handler.AuthMe)
	subRoute.Delete("/logout/", m.handler.AuthLogout)
}
