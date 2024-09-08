package datastruct

import "time"

type ExampleList struct {
	UUID      string    `db:"uuid" json:"id"`
	Name      any       `db:"nama" json:"name"`
	Detail    any       `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	TotalRows int       `db:"total_rows" json:"-"`
}

type ExampleDetail struct {
	UUID      string    `db:"uuid" json:"id"`
	Name      any       `db:"nama" json:"name"`
	Detail    any       `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
