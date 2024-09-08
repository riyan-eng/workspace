package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Reset Password
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.PerangkatResetPassword	true  "body"
// @Router      /perangkat/{id}/reset-password/ [patch]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatResetPassword(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	payload := new(dto.PerangkatResetPassword)

	if err := c.BodyParser(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)

	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)

	}

	if err := m.perangkatService.ResetPassword(&ctx, &entity.ServPerangkatResetPassword{
		Id:       &id,
		Password: &payload.Password,
	}); err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	data := map[string]any{
		"id": id,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_UPDATE"), 200)
}
