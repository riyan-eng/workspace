package router

func (m *routeStruct) Object() {
	subRoute := m.app.Group("/object")
	subRoute.GET("/:id/:name", m.handler.ObjectView)
	subRoute.POST("/", m.handler.ObjectUpload)
	subRoute.DELETE("/:id/", m.handler.ObjectRemove)
}
