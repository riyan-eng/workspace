package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Logout
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Router		/auth/logout/ [delete]
// @Security	ApiKeyAuth
func (m *ServiceServer) AuthLogout(c *fiber.Ctx) error {
	ctx := context.Background()
	user := util.CurrentUser(c)

	if err := m.authService.Logout(&ctx, &entity.ServAuthLogout{UserId: &user.UserId}); err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
	}

	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_LOGOUT"))
}
