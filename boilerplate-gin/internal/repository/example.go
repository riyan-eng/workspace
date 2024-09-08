package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/infrastructure"
	"server/internal/datastruct"
	"server/util"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type ExampleRepository interface {
	List(ctx *context.Context, limit *int, offset *int) (*[]datastruct.ExampleList, *int, *util.Error)
	Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error)
}

type exampleRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
}

func (m *exampleRepository) List(ctx *context.Context, limit *int, offset *int) (*[]datastruct.ExampleList, *int, *util.Error) {
	data := make([]datastruct.ExampleList, 0)
	countRow := new(int)

	query := fmt.Sprintf(`
	select uuid, nama, detail, created_at, updated_at, count(uuid) over() as total_rows from example 
	where lower(nama) like lower('%%%v%%') order by nama %v limit %v offset %v
	`, "", "desc", *limit, *offset)
	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return &data, countRow, &util.Error{
			Errors: err.Error(),
		}
	}

	if err := scan.Rows(&data, sqlRows); err != nil {
		return &data, countRow, &util.Error{
			Errors: err.Error(),
		}
	}

	for _, d := range data {
		countRow = &d.TotalRows
		break
	}

	return &data, countRow, &util.Error{}
}

func (m *exampleRepository) Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error) {
	data := new(datastruct.ExampleDetail)

	query := fmt.Sprintf(`
	select uuid, nama, detail, created_at, updated_at from example 
	where uuid = '%v'
	`, *id)
	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return data, &util.Error{
			Errors: err.Error(),
		}
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, &util.Error{
			Errors:     err.Error(),
			StatusCode: 400,
			Message:    infrastructure.Localize("NOT_FOUND"),
		}
	}

	return data, &util.Error{}
}
