package api

import (
	"context"
	"os"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Remove
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /object/{id}/ [delete]
// @Security ApiKeyAuth
func (m *ServiceServer) ObjectRemove(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	data, err := m.objectService.Detail(&ctx, &entity.ServObjectDetail{Id: &id})
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	go func(path string) {
		os.Remove(path)
	}(data.Path)

	rdata := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(rdata, nil, infrastructure.Localize("OK_DELETE"))
}
