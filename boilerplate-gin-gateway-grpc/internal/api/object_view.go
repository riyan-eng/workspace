package api

import (
	"context"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     View
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       name	path	string	true	"fill with random"
// @Router      /object/{id}/{name} [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ObjectView(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	data, err := m.objectService.Detail(&ctx, &entity.ServObjectDetail{Id: &id})
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	c.Header("Content-Type", data.ContentType)
	c.FileAttachment(data.Path, data.Name)
}
