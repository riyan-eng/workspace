package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Detail
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /perangkat/{id}/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatDetail(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")

	data, err := m.perangkatService.Detail(&ctx, &entity.ServPerangkatDetail{Id: &id})
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
