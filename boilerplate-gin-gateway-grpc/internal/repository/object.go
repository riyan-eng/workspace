package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/infrastructure"
	"server/internal/datastruct"
	"server/internal/model"
	"server/util"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type ObjectRepository interface {
	Create(ctx *context.Context, mdl *model.Object) *util.Error
	Detail(ctx *context.Context, id *string) (*datastruct.ObjectDetail, *util.Error)
	Delete(ctx *context.Context, id *string) *util.Error
}

type objectRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
}

func (m *objectRepository) Create(ctx *context.Context, mdl *model.Object) *util.Error {
	sqlRslt, errT := m.sqlxDB.NamedExecContext(*ctx, `
	insert into objects (uuid, name, owner, size, content_type, path, url) 
	values (:id, :name, :owner, :size, :content_type, :path, :url)
	`, mdl)
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     "no rows",
			Message:    infrastructure.Localize("FAILED_CREATE_NO_DATA"),
			StatusCode: 400,
		}
	}

	return &util.Error{}
}

func (m *objectRepository) Detail(ctx *context.Context, id *string) (*datastruct.ObjectDetail, *util.Error) {
	data := new(datastruct.ObjectDetail)

	query := fmt.Sprintf(`
	select o."uuid", o."name", o."size", o.content_type, o."path", o.url from objects o 
	where o."uuid" = '%v'
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

func (m *objectRepository) Delete(ctx *context.Context, id *string) *util.Error {
	sqlRslt, err := m.sqlxDB.ExecContext(*ctx, fmt.Sprintf("delete from objects where uuid = '%v'", *id))
	if err != nil {
		return &util.Error{
			Errors: err.Error(),
		}
	}

	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     "no data",
			Message:    infrastructure.Localize("FAILED_DELETE_NO_DATA"),
			StatusCode: 400,
		}
	}

	return &util.Error{}
}
