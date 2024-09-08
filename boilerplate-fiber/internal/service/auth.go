package service

import (
	"context"
	"server/config"
	"server/internal/datastruct"
	"server/internal/entity"
	"server/internal/repository"
	"server/util"
)

type AuthService interface {
	Login(ctx *context.Context, ent *entity.ServAuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error)
	Refresh(ctx *context.Context, ent *entity.ServAuthRefresh) (*datastruct.AuthToken, *util.Error)
	Me(ctx *context.Context, ent *entity.ServAuthMe) (*datastruct.AuthMe, *util.Error)
	Logout(ctx *context.Context, ent *entity.ServAuthLogout) *util.Error
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao *repository.DAO) AuthService {
	return &authService{
		dao: *dao,
	}
}

func (m *authService) Login(ctx *context.Context, ent *entity.ServAuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error) {
	token := new(datastruct.AuthToken)

	data, err := m.dao.NewAuthRepository().Login(ctx, ent.Username)
	if err.Errors != nil {
		return data, token, err
	}

	if !data.IsActive {
		return data, token, &util.Error{
			Errors:     "user",
			Message:    "user tidak aktif",
			StatusCode: 400,
		}
	}

	// verify password
	if !util.VerifyHash(data.Password, *ent.Password) {
		return data, token, &util.Error{
			Errors:     "password",
			Message:    "Username atau Password yang anda inputkan salah, silakan menginputkan username dan password yang benar",
			StatusCode: 400,
		}
	}

	accessToken, accessExpire, errT := util.NewToken().CreateAccess(ctx, &data.Id, &data.RoleCode)
	if errT != nil {
		return data, token, &util.Error{
			Errors: errT.Error(),
		}
	}
	refreshToken, refreshExpired, errT := util.NewToken().CreateRefresh(ctx, &data.Id, &data.RoleCode)
	if errT != nil {
		return data, token, &util.Error{
			Errors: errT.Error(),
		}
	}
	enforce := config.NewEnforcer()
	enforce.AddRoleForUser(data.Id, data.RoleCode)
	return data, &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Refresh(ctx *context.Context, ent *entity.ServAuthRefresh) (*datastruct.AuthToken, *util.Error) {
	newRefresh := new(datastruct.AuthToken)
	claim, errT := util.NewToken().ParseRefresh(ent.Token)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT.Error(),
			StatusCode: 401,
		}
	}

	if errT := util.NewToken().ValidateRefresh(ctx, claim); errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT.Error(),
			StatusCode: 401,
		}
	}

	accessToken, accessExpire, errT := util.NewToken().CreateAccess(ctx, &claim.UserId, &claim.RoleCode)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors: errT.Error(),
		}
	}
	refreshToken, refreshExpired, errT := util.NewToken().CreateRefresh(ctx, &claim.UserId, &claim.RoleCode)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors: errT.Error(),
		}
	}

	return &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Logout(ctx *context.Context, ent *entity.ServAuthLogout) *util.Error {
	err := m.dao.NewAuthRepository().Logout(ctx, ent.UserId)
	if err.Errors != nil {
		return err
	}

	return &util.Error{}
}

func (m *authService) Me(ctx *context.Context, ent *entity.ServAuthMe) (*datastruct.AuthMe, *util.Error) {
	data, err := m.dao.NewAuthRepository().Me(ctx, ent.UserId)
	if err.Errors != nil {
		return data, err
	}

	return data, &util.Error{}
}
