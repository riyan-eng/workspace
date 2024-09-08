package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthLogin	true  "body"
// @Router		/auth/login/ [post]
func (m *ServiceServer) AuthLogin(c *fiber.Ctx) error {
	ctx := context.Background()
	payload := new(dto.AuthLogin)

	if err := c.BodyParser(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
	}
	user, token, err := m.authService.Login(&ctx, &entity.ServAuthLogin{
		Username: &payload.Username,
		Password: &payload.Password,
	})

	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
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
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
