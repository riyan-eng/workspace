package router

func (m *routeStruct) Index() {
	subRoute := m.app.Group("/")
	subRoute.Get("/", m.handler.Index)
}
