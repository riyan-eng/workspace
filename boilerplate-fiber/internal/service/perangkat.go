package service

import (
	"context"
	"database/sql"
	"server/internal/datastruct"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/util"

	"github.com/google/uuid"
)

type PerangkatService interface {
	List(ctx *context.Context, ent *entity.ServPerangkatList) (*[]datastruct.PerangkatList, *int, *util.Error)
	Create(ctx *context.Context, ent *entity.ServPerangkatCreate) *util.Error
	Patch(ctx *context.Context, ent *entity.ServPerangkatPatch) *util.Error
	ResetPassword(ctx *context.Context, ent *entity.ServPerangkatResetPassword) *util.Error
	Detail(ctx *context.Context, ent *entity.ServPerangkatDetail) (*datastruct.PerangkatDetail, *util.Error)
	Delete(ctx *context.Context, ent *entity.ServPerangkatDelete) *util.Error
}

type perangkatService struct {
	dao repository.DAO
}

func NewPerangkatService(dao *repository.DAO) PerangkatService {
	return &perangkatService{
		dao: *dao,
	}
}

func (m *perangkatService) List(ctx *context.Context, ent *entity.ServPerangkatList) (*[]datastruct.PerangkatList, *int, *util.Error) {
	data, countRow, err := m.dao.NewPerangkatRepository().List(ctx, ent.Search, ent.Limit, ent.Offset)
	if err.Errors != nil {
		// custom err
		return data, countRow, err
	}

	return data, countRow, &util.Error{}
}

func (m *perangkatService) Detail(ctx *context.Context, ent *entity.ServPerangkatDetail) (*datastruct.PerangkatDetail, *util.Error) {
	data, err := m.dao.NewPerangkatRepository().Detail(ctx, ent.Id)
	if err.Errors != nil {
		// custom err
		return data, err
	}

	return data, &util.Error{}
}

func (m *perangkatService) Delete(ctx *context.Context, ent *entity.ServPerangkatDelete) *util.Error {
	if err := m.dao.NewPerangkatRepository().Delete(ctx, ent.Id); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *perangkatService) Create(ctx *context.Context, ent *entity.ServPerangkatCreate) *util.Error {
	hashPassword, errT := util.GenerateHash(ent.Password)
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}

	modelUser := model.User{
		Id:         *ent.Id,
		Username:   sql.NullString{String: *ent.Username, Valid: util.NewIsValid().String(ent.Username)},
		Password:   sql.NullString{String: hashPassword, Valid: true},
		BirthPlace: sql.NullString{String: *ent.BirthPlace, Valid: util.NewIsValid().String(ent.BirthPlace)},
		BirthDate:  sql.NullString{String: *ent.BirthDate, Valid: util.NewIsValid().String(ent.BirthDate)},
		Address:    sql.NullString{String: *ent.Address, Valid: util.NewIsValid().String(ent.Address)},
		PhotoUrl:   sql.NullString{String: *ent.PhotoUrl, Valid: util.NewIsValid().String(ent.PhotoUrl)},
		IsActive:   sql.NullBool{Bool: true, Valid: true},
	}

	modelUserData := model.UserData{
		Id:          uuid.NewString(),
		UserId:      sql.NullString{String: modelUser.Id, Valid: true},
		RoleCode:    sql.NullString{String: *ent.RoleCode, Valid: true},
		JabatanCode: sql.NullString{String: *ent.JabatanCode, Valid: true},
	}
	if err := m.dao.NewPerangkatRepository().Create(ctx, &modelUser, &modelUserData); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *perangkatService) Patch(ctx *context.Context, ent *entity.ServPerangkatPatch) *util.Error {
	modelUser := model.User{
		Id:         *ent.Id,
		Username:   sql.NullString{String: *ent.Username, Valid: util.NewIsValid().String(ent.Username)},
		BirthPlace: sql.NullString{String: *ent.BirthPlace, Valid: util.NewIsValid().String(ent.BirthPlace)},
		BirthDate:  sql.NullString{String: *ent.BirthDate, Valid: util.NewIsValid().String(ent.BirthDate)},
		Address:    sql.NullString{String: *ent.Address, Valid: util.NewIsValid().String(ent.Address)},
		PhotoUrl:   sql.NullString{String: *ent.PhotoUrl, Valid: util.NewIsValid().String(ent.PhotoUrl)},
	}

	modelUserData := model.UserData{
		Id:          uuid.NewString(),
		UserId:      sql.NullString{String: modelUser.Id, Valid: true},
		RoleCode:    sql.NullString{String: *ent.RoleCode, Valid: true},
		JabatanCode: sql.NullString{String: *ent.JabatanCode, Valid: true},
	}
	if err := m.dao.NewPerangkatRepository().Patch(ctx, &modelUser, &modelUserData); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *perangkatService) ResetPassword(ctx *context.Context, ent *entity.ServPerangkatResetPassword) *util.Error {
	hashPassword, errT := util.GenerateHash(ent.Password)
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	modelUser := model.User{
		Id:       *ent.Id,
		Password: sql.NullString{String: hashPassword, Valid: true},
	}

	modelUserData := model.UserData{
		Id:     uuid.NewString(),
		UserId: sql.NullString{String: modelUser.Id, Valid: true},
	}
	if err := m.dao.NewPerangkatRepository().Patch(ctx, &modelUser, &modelUserData); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}
