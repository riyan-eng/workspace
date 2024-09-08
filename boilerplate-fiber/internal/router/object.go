package router

import "server/internal/middleware"

func (m *routeStruct) Object() {
	subRoute := m.app.Group("/object")
	subRoute.Get("/:id/:name", m.handler.ObjectView)
	subRoute.Use(middleware.Jwt())
	subRoute.Post("/", m.handler.ObjectUpload)
	subRoute.Delete("/:id/", m.handler.ObjectRemove)
}
