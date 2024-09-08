package api

import (
	"context"
	"server/infrastructure"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /example/{id}/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleDetail(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")

	data, err := m.exampleService.Detail(&ctx, &id)
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
