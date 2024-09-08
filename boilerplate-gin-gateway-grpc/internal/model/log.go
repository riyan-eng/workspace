package model

import "database/sql"

type Log struct {
	Id         string         `db:"id"`
	Path       string         `db:"path"`
	Method     string         `db:"method"`
	StatusCode int            `db:"status_code"`
	Interval   float64        `db:"interval"`
	Body       sql.NullString `db:"body"`
	UserId     sql.NullString `db:"user_id"`
}
