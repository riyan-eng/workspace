package router

import "server/internal/middleware"

func (m *routeStruct) Export() {
	subRoute := m.app.Group("/export")
	subRoute.Use(middleware.Jwt())
	subRoute.Use(middleware.Permission())
}
