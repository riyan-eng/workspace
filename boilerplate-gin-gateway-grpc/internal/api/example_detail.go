package api

import (
	"context"
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /example/{id} [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleDetail(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	res, err := m.exampleRpcServer.Detail(ctx, &pb.TaskDetailRequest{
		Id: id,
	})
	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	util.NewResponse(c).Success(res.Data, nil, infrastructure.Localize("OK_READ"))
}
