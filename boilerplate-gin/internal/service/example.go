package service

import (
	"context"
	"server/internal/datastruct"
	"server/internal/repository"
	"server/util"
)

type ExampleService interface {
	List(ctx *context.Context, limit, offset *int) (*[]datastruct.ExampleList, *int, *util.Error)
	Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error)
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao *repository.DAO) ExampleService {
	return &exampleService{
		dao: *dao,
	}
}

func (m *exampleService) List(ctx *context.Context, limit, offset *int) (*[]datastruct.ExampleList, *int, *util.Error) {
	data, countRow, err := m.dao.NewExampleRepository().List(ctx, limit, offset)
	if err.Errors != nil {
		// custom err
		return data, countRow, err
	}

	return data, countRow, &util.Error{}
}

func (m *exampleService) Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error) {
	data, err := m.dao.NewExampleRepository().Detail(ctx, id)
	if err.Errors != nil {
		// custom err
		return data, err
	}

	return data, &util.Error{}
}
