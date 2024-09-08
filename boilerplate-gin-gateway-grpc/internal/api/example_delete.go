package api

import (
	"context"
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Delete
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /example/{id} [delete]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleDelete(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	_, err := m.exampleRpcServer.Delete(ctx, &pb.TaskDeleteRequest{
		Id: id,
	})
	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
