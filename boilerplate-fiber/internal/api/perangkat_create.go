package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary     Create
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       body	body  dto.PerangkatCreate	true  "body"
// @Router		/perangkat/ [post]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatCreate(c *fiber.Ctx) error {
	ctx := context.Background()
	payload := new(dto.PerangkatCreate)

	if err := c.BodyParser(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)

	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)

	}

	id := uuid.NewString()

	if errT := util.NewHelper().CheckExistJabatan(&id, &payload.JabatanCode); errT != nil {
		return util.NewResponse(c).Error(errT.Error(), "jabatan sudah digunakan perangkat lain", 400)

	}

	password := strings.ReplaceAll(payload.BirthDate, "-", "")
	roleCode := util.NewEnum().JabatanRole(&payload.JabatanCode)

	if err := m.perangkatService.Create(&ctx, &entity.ServPerangkatCreate{
		Id:          &id,
		Username:    &payload.Username,
		Password:    &password,
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
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"), 201)
}
