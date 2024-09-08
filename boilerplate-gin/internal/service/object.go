package service

import (
	"context"
	"database/sql"
	"server/internal/datastruct"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/util"
)

type ObjectService interface {
	Create(ctx *context.Context, ent *entity.ServObjectCreate) *util.Error
	Detail(ctx *context.Context, ent *entity.ServObjectDetail) (*datastruct.ObjectDetail, *util.Error)
	Delete(ctx *context.Context, ent *entity.ServObjectDelete) *util.Error
}

type objectService struct {
	dao repository.DAO
}

func NewObjectService(dao *repository.DAO) ObjectService {
	return &objectService{
		dao: *dao,
	}
}

func (m *objectService) Detail(ctx *context.Context, ent *entity.ServObjectDetail) (*datastruct.ObjectDetail, *util.Error) {
	data, err := m.dao.NewObjectRepository().Detail(ctx, ent.Id)
	if err.Errors != nil {
		// custom err
		return data, err
	}

	return data, &util.Error{}
}

func (m *objectService) Delete(ctx *context.Context, ent *entity.ServObjectDelete) *util.Error {
	if err := m.dao.NewObjectRepository().Delete(ctx, ent.Id); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *objectService) Create(ctx *context.Context, ent *entity.ServObjectCreate) *util.Error {
	mdl := model.Object{
		Id:          *ent.Id,
		Name:        sql.NullString{String: *ent.Name, Valid: util.NewIsValid().String(ent.Name)},
		Owner:       sql.NullString{String: *ent.Owner, Valid: util.NewIsValid().String(ent.Owner)},
		Size:        sql.NullInt64{Int64: int64(*ent.Size), Valid: util.NewIsValid().Int(ent.Size)},
		ContentType: sql.NullString{String: *ent.ContentType, Valid: util.NewIsValid().String(ent.ContentType)},
		Url:         sql.NullString{String: *ent.Url, Valid: util.NewIsValid().String(ent.Url)},
		Path:        sql.NullString{String: *ent.Path, Valid: util.NewIsValid().String(ent.Path)},
	}
	if err := m.dao.NewObjectRepository().Create(ctx, &mdl); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}
