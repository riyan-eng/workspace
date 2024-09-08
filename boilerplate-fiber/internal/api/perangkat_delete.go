package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Delete
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /perangkat/{id}/ [delete]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatDelete(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")

	if err := m.perangkatService.Delete(&ctx, &entity.ServPerangkatDelete{Id: &id}); err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	data := map[string]any{
		"id": id,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_DELETE"))
}
