package api

import (
	"context"
	"server/infrastructure"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /example/{id}/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleDetail(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	data, err := m.exampleService.Detail(&ctx, &id)
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
