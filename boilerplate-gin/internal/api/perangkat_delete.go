package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Delete
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /perangkat/{id}/ [delete]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatDelete(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	if err := m.perangkatService.Delete(&ctx, &entity.ServPerangkatDelete{Id: &id}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_DELETE"))
}
