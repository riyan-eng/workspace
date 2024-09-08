package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Create
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       body	body  dto.PerangkatCreate	true  "body"
// @Router		/perangkat/ [post]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatCreate(c *gin.Context) {
	ctx := context.Background()
	payload := new(dto.PerangkatCreate)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	id := uuid.NewString()

	if errT := util.NewHelper().CheckExistJabatan(&id, &payload.JabatanCode); errT != nil {
		util.NewResponse(c).Error(errT.Error(), "jabatan sudah digunakan perangkat lain", 400)
		return
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
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"), 201)
}
