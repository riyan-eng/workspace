package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Reset Password
// @Tags       	Perangkat
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.PerangkatResetPassword	true  "body"
// @Router      /perangkat/{id}/reset-password/ [patch]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatResetPassword(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	payload := new(dto.PerangkatResetPassword)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	if err := m.perangkatService.ResetPassword(&ctx, &entity.ServPerangkatResetPassword{
		Id:       &id,
		Password: &payload.Password,
	}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_UPDATE"), 200)
}
