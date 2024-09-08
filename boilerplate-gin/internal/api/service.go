package api

import "server/internal/service"

type ServiceServer struct {
	exampleService   service.ExampleService
	authService      service.AuthService
	perangkatService service.PerangkatService
	objectService    service.ObjectService
}

func NewService(
	exampleService service.ExampleService,
	authService service.AuthService,
	perangkatService service.PerangkatService,
	objectService service.ObjectService,
) *ServiceServer {
	return &ServiceServer{
		exampleService:   exampleService,
		authService:      authService,
		perangkatService: perangkatService,
		objectService:    objectService,
	}
}
