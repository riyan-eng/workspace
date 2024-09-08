package api

import (
	"server/internal/service"
	"server/pb"
)

type ServiceServer struct {
	exampleService   service.ExampleService
	authService      service.AuthService
	perangkatService service.PerangkatService
	objectService    service.ObjectService
	exampleRpcServer pb.TaskServiceClient
	authRpcServer    pb.AuthServiceClient
}

func NewService(
	exampleService service.ExampleService,
	authService service.AuthService,
	perangkatService service.PerangkatService,
	objectService service.ObjectService,
	exampleRpcServer pb.TaskServiceClient,
	authRpcServer pb.AuthServiceClient,
) *ServiceServer {
	return &ServiceServer{
		exampleService:   exampleService,
		authService:      authService,
		perangkatService: perangkatService,
		objectService:    objectService,
		exampleRpcServer: exampleRpcServer,
		authRpcServer:    authRpcServer,
	}
}
