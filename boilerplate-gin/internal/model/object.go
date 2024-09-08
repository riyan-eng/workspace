package model

import (
	"database/sql"
)

type Object struct {
	Id          string         `db:"id"`
	Name        sql.NullString `db:"name"`
	Owner       sql.NullString `db:"owner"`
	Size        sql.NullInt64  `db:"size"`
	ContentType sql.NullString `db:"content_type"`
	Url         sql.NullString `db:"url"`
	Path        sql.NullString `db:"path"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}
