package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthLogin	true  "body"
// @Router		/auth/login/ [post]
func (m *ServiceServer) AuthLogin(c *gin.Context) {
	ctx := context.Background()
	payload := new(dto.AuthLogin)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}
	user, token, err := m.authService.Login(&ctx, &entity.ServAuthLogin{
		Username: &payload.Username,
		Password: &payload.Password,
	})

	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"access_token":    token.AccessToken,
		"access_expired":  token.AccessExpired.Time.Local(),
		"refresh_token":   token.RefreshToken,
		"refresh_expired": token.RefreshExpired.Time.Local(),
		"user": map[string]any{
			"username":     user.Username,
			"jabatan_name": user.JabatanName,
		},
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
