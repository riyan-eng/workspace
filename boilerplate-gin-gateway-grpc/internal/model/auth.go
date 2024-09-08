package model

import (
	"database/sql"
)

type User struct {
	Id         string         `db:"id"`
	Username   sql.NullString `db:"username"`
	Password   sql.NullString `db:"password"`
	BirthPlace sql.NullString `db:"birth_place"`
	BirthDate  sql.NullString `db:"birth_date"`
	Address    sql.NullString `db:"address"`
	PhotoUrl   sql.NullString `db:"photo_url"`
	IsActive   sql.NullBool   `db:"is_active"`
	CreatedAt  sql.NullTime   `db:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at"`
}

type UserData struct {
	Id          string         `db:"id"`
	UserId      sql.NullString `db:"user_id"`
	RoleCode    sql.NullString `db:"role_code"`
	JabatanCode sql.NullString `db:"jabatan_code"`
}
