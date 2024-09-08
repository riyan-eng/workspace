package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Me
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Router		/auth/me/ [get]
// @Security	ApiKeyAuth
func (m *ServiceServer) AuthMe(c *fiber.Ctx) error {
	ctx := context.Background()
	user := util.CurrentUser(c)

	data, err := m.authService.Me(&ctx, &entity.ServAuthMe{UserId: &user.UserId})
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
