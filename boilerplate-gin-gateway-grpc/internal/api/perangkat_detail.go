package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Detail
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /perangkat/{id}/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatDetail(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	data, err := m.perangkatService.Detail(&ctx, &entity.ServPerangkatDetail{Id: &id})
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
