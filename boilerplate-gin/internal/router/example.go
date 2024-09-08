package router

func (m *routeStruct) Example() {
	subRoute := m.app.Group("/example")
	subRoute.GET("/", m.handler.ExampleList)
	subRoute.GET("/:id/", m.handler.ExampleDetail)
}
