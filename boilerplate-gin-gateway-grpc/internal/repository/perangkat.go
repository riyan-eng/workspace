package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/infrastructure"
	"server/internal/datastruct"
	"server/internal/model"
	"server/util"
	"strings"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type PerangkatRepository interface {
	List(ctx *context.Context, search *string, limit, offset *int) (*[]datastruct.PerangkatList, *int, *util.Error)
	Create(ctx *context.Context, modelUser *model.User, modelUserData *model.UserData) *util.Error
	Patch(ctx *context.Context, modelUser *model.User, modelUserData *model.UserData) *util.Error
	Detail(ctx *context.Context, id *string) (*datastruct.PerangkatDetail, *util.Error)
	Delete(ctx *context.Context, id *string) *util.Error
}

type perangkatRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
}

func (m *perangkatRepository) Create(ctx *context.Context, modelUser *model.User, modelUserData *model.UserData) *util.Error {
	tx := m.sqlxDB.MustBegin()
	_, errT := tx.NamedExecContext(*ctx, `
	insert into users (uuid, username, password, birth_place, birth_date, address, photo_url, is_active) values (:id, :username, :password, :birth_place, :birth_date, :address, :photo_url, :is_active)
	`, modelUser)
	if errT != nil {
		if strings.Contains(errT.Error(), `duplicate key value violates unique constraint "users_username_is_delete_key"`) {
			return &util.Error{
				Errors:     "duplicate",
				Message:    "username telah digunakan.",
				StatusCode: 409,
			}
		}

		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_CREATE_NO_DATA"),
		}
	}
	_, errT = tx.NamedExecContext(*ctx, `
	insert into user_datas (uuid, user_uuid, role_code, jabatan_code) values (:id, :user_id, :role_code, :jabatan_code)
	`, modelUserData)

	if errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_CREATE_NO_DATA"),
		}
	}
	if errT := tx.Commit(); errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_CREATE_NO_DATA"),
		}
	}
	return &util.Error{}
}

func (m *perangkatRepository) Patch(ctx *context.Context, modelUser *model.User, modelUserData *model.UserData) *util.Error {
	tx := m.sqlxDB.MustBegin()
	_, errT := tx.NamedExecContext(*ctx, `
	update users set username=coalesce(:username, username), password=coalesce(:password, password), birth_place=coalesce(:birth_place, birth_place), birth_date=coalesce(:birth_date, birth_date), 
	address=coalesce(:address, address), photo_url=coalesce(:photo_url, photo_url), is_active=coalesce(:is_active, is_active), 
	updated_at=now() where uuid=:id
	`, modelUser)
	if errT != nil {
		if strings.Contains(errT.Error(), `duplicate key value violates unique constraint "users_username_is_delete_key"`) {
			return &util.Error{
				Errors:     "duplicate",
				Message:    "username telah digunakan.",
				StatusCode: 409,
			}
		}
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
		}
	}
	_, errT = tx.NamedExecContext(*ctx, `
	update user_datas set role_code=coalesce(:role_code, role_code), jabatan_code=coalesce(:jabatan_code, jabatan_code), updated_at=now() where user_uuid=:user_id
	`, modelUserData)

	if errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
		}
	}
	if errT := tx.Commit(); errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
		}
	}
	return &util.Error{}
}

func (m *perangkatRepository) List(ctx *context.Context, search *string, limit, offset *int) (*[]datastruct.PerangkatList, *int, *util.Error) {
	data := make([]datastruct.PerangkatList, 0)
	countRow := new(int)

	query := fmt.Sprintf(`
	select u.id, u."uuid", u.username, ud.jabatan_code, j.name as jabatan_name, u.birth_place, u.birth_date::text, u.address, u.photo_url, 
	count(u.uuid) over() as total_rows
	from users u 
	left join user_datas ud on ud.user_uuid = u.uuid 
	left join jabatan j on ud.jabatan_code = j.code 
	where ud.role_code not in ('ADMIN') and lower(u.username) like lower('%%%v%%') order by u.created_at %v limit %v offset %v
	`, *search, "desc", *limit, *offset)
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

	concurrency := 10
	sem := make(chan bool, concurrency)
	for i, d := range data {
		sem <- true
		go func(i int) {
			defer func() { <-sem }()
			data[i].Idf = fmt.Sprintf("%03d", d.Id)
		}(i)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	for _, d := range data {
		countRow = &d.TotalRows
		break
	}

	return &data, countRow, &util.Error{}
}

func (m *perangkatRepository) Detail(ctx *context.Context, id *string) (*datastruct.PerangkatDetail, *util.Error) {
	data := new(datastruct.PerangkatDetail)

	query := fmt.Sprintf(`
	select u.id, u."uuid", u.username, ud.jabatan_code, j.name as jabatan_name, u.birth_place, u.birth_date, u.address, u.photo_url 
	from users u 
	left join user_datas ud on ud.user_uuid = u.uuid 
	left join jabatan j on ud.jabatan_code = j.code 
	where u.uuid = '%v'
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

	data.Idf = fmt.Sprintf("%03d", data.Id)

	return data, &util.Error{}
}

func (m *perangkatRepository) Delete(ctx *context.Context, id *string) *util.Error {
	sqlRslt, err := m.sqlxDB.ExecContext(*ctx, fmt.Sprintf("delete from users where uuid = '%v'", *id))
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
