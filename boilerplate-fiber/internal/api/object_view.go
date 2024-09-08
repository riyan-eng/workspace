package api

import (
	"context"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     View
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       name	path	string	true	"fill with random"
// @Router      /object/{id}/{name} [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ObjectView(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")

	data, err := m.objectService.Detail(&ctx, &entity.ServObjectDetail{Id: &id})
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	c.Response().Header.SetContentType(data.ContentType)
	return c.SendFile(data.Path)
}
