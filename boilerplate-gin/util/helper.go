package util

import (
	"database/sql"
	"fmt"
	"server/infrastructure"

	"github.com/blockloop/scan/v2"
)

type helper struct{}

func NewHelper() *helper {
	return &helper{}
}

func (m *helper) CheckExistJabatan(userId, jabatanCode *string) error {
	if *jabatanCode == "Q" {
		return nil
	}

	var id string
	query := fmt.Sprintf(`
	select u."uuid"
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	where ud.jabatan_code = '%s'
	limit 1
	`, *jabatanCode)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		return err
	}
	err = scan.Row(&id, sqlrows)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return fmt.Errorf("error query scan")
	}
	if *userId == id {
		return nil
	}
	return fmt.Errorf("user jabatan existing")
}
