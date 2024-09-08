package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.ExamplePut	true  "body"
// @Router      /example/{id} [put]
// @Security ApiKeyAuth
func (m *ServiceServer) ExamplePut(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	payload := new(dto.ExamplePut)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	_, err := m.exampleRpcServer.Put(ctx, &pb.TaskPutRequest{
		Id:     id,
		Name:   payload.Name,
		Detail: payload.Detail,
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
