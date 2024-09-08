package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Patch
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.PerangkatPatch	true  "body"
// @Router      /perangkat/{id}/ [patch]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatPatch(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	payload := new(dto.PerangkatPatch)

	if err := c.BodyParser(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)

	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)

	}

	if errT := util.NewHelper().CheckExistJabatan(&id, &payload.JabatanCode); errT != nil {
		return util.NewResponse(c).Error(errT.Error(), "jabatan sudah digunakan perangkat lain", 400)

	}

	roleCode := util.NewEnum().JabatanRole(&payload.JabatanCode)

	if err := m.perangkatService.Patch(&ctx, &entity.ServPerangkatPatch{
		Id:          &id,
		Username:    &payload.Username,
		RoleCode:    &roleCode,
		BirthPlace:  &payload.BirthPlace,
		BirthDate:   &payload.BirthDate,
		JabatanCode: &payload.JabatanCode,
		Address:     &payload.Address,
		PhotoUrl:    &payload.PhotoUrl,
	}); err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	data := map[string]any{
		"id": id,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_UPDATE"), 200)
}
