package router

func (m *routeStruct) Index() {
	subRoute := m.app.Group("/")
	subRoute.GET("/", m.handler.Index)
}
