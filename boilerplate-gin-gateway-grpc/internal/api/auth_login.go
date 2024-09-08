package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
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
	res, err := m.authRpcServer.Login(ctx, &pb.AuthLoginRequest{
		Username: payload.Username,
		Password: payload.Password,
	})

	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	data := map[string]any{
		"access_token":    res.AccessToken,
		"access_expired":  res.AccessExpired,
		"refresh_token":   res.RefreshToken,
		"refresh_expired": res.RefreshExpired,
		"user": map[string]any{
			"username": res.Username,
		},
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
