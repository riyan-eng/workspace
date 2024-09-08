package router

func (m *routeStruct) Example() {
	subRoute := m.app.Group("/example")
	subRoute.POST("", m.handler.ExampleCreate)
	subRoute.GET("", m.handler.ExampleList)
	subRoute.GET("/:id", m.handler.ExampleDetail)
	subRoute.PUT("/:id", m.handler.ExamplePut)
	subRoute.DELETE("/:id", m.handler.ExampleDelete)
}
