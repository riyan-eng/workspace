package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Logout
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Router		/auth/logout/ [delete]
// @Security	ApiKeyAuth
func (m *ServiceServer) AuthLogout(c *gin.Context) {
	ctx := context.Background()
	user := util.CurrentUser(c)

	if err := m.authService.Logout(&ctx, &entity.ServAuthLogout{UserId: &user.UserId}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_LOGOUT"))
}
