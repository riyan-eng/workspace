package api

import (
	"context"
	"os"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Remove
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /object/{id}/ [delete]
// @Security ApiKeyAuth
func (m *ServiceServer) ObjectRemove(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")

	data, err := m.objectService.Detail(&ctx, &entity.ServObjectDetail{Id: &id})
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	go func(path string) {
		os.Remove(path)
	}(data.Path)

	rdata := map[string]any{
		"id": id,
	}
	return util.NewResponse(c).Success(rdata, nil, infrastructure.Localize("OK_DELETE"))
}
