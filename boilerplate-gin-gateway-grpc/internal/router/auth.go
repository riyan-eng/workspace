package router

import (
	"server/internal/middleware"
)

func (m *routeStruct) Authentication() {
	subRoute := m.app.Group("/auth")
	subRoute.POST("/login/", m.handler.AuthLogin)
	subRoute.POST("/refresh/", m.handler.AuthRefresh)
	subRoute.Use(middleware.Jwt())
	subRoute.GET("/me/", m.handler.AuthMe)
	subRoute.DELETE("/logout/", m.handler.AuthLogout)
}
