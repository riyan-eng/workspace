package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Patch
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.PerangkatPatch	true  "body"
// @Router      /perangkat/{id}/ [patch]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatPatch(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	payload := new(dto.PerangkatPatch)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	if errT := util.NewHelper().CheckExistJabatan(&id, &payload.JabatanCode); errT != nil {
		util.NewResponse(c).Error(errT.Error(), "jabatan sudah digunakan perangkat lain", 400)
		return
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
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_UPDATE"), 200)
}
