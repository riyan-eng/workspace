package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/datastruct"
	"server/util"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthRepository interface {
	Login(ctx *context.Context, username *string) (*datastruct.AuthLoginData, *util.Error)
	Logout(ctx *context.Context, userId *string) *util.Error
	Me(ctx *context.Context, userId *string) (*datastruct.AuthMe, *util.Error)
}

type authRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
	cache  *redis.Client
}

func (m *authRepository) Login(ctx *context.Context, username *string) (*datastruct.AuthLoginData, *util.Error) {
	data := new(datastruct.AuthLoginData)

	query := fmt.Sprintf(`
	select u."uuid", u.username, u."password", u.is_active, j.name as jabatan_name, r.code as role_code, r."name" as role_name 
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	left join roles r on r.code = ud.role_code  
	left join jabatan j on ud.jabatan_code = j.code 
	where u.username = '%v' and u.is_delete = false
	limit 1
	`, *username)

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
			Message:    "User belum terdaftar",
		}
	}

	return data, &util.Error{}
}

func (m *authRepository) Logout(ctx *context.Context, userId *string) *util.Error {
	if err := m.cache.Del(*ctx, fmt.Sprintf("access-token-%s", *userId)).Err(); err != nil {
		return &util.Error{
			Errors: err.Error(),
		}
	}
	return &util.Error{}
}

func (m *authRepository) Me(ctx *context.Context, userId *string) (*datastruct.AuthMe, *util.Error) {
	data := new(datastruct.AuthMe)

	query := fmt.Sprintf(`
	select u."uuid", u.username, r.code as role_code, r."name" as role_name, ud.jabatan_code, j.name as jabatan_name, u.birth_place, u.birth_date::text, u.address, u.photo_url 
	from users u 
	left join user_datas ud on ud.user_uuid = u.uuid 
	left join roles r on r.code = ud.role_code  
	left join jabatan j on ud.jabatan_code = j.code 
	where u.uuid = '%v' 
	limit 1
	`, *userId)

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
			Message:    "tidak terdaftar",
		}
	}

	return data, &util.Error{}
}
