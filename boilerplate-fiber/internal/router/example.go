package router

func (m *routeStruct) Example() {
	subRoute := m.app.Group("/example")
	subRoute.Get("/", m.handler.ExampleList)
	subRoute.Get("/:id/", m.handler.ExampleDetail)
}
