package util

import (
	"fmt"
	"server/infrastructure"
	"server/internal/datastruct"

	"github.com/blockloop/scan/v2"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type queryStruct struct{}

func NewQuery() *queryStruct {
	return &queryStruct{}
}

func (q *queryStruct) PekerjaanListName() []string {
	names := make([]string, 0)
	query := `
	select p.name
	from pekerjaan p
	order by p.created_at asc
	`
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Rows(&names, sqlrows); err != nil {
		hlog.Info(err)
		return names

	}
	return names
}

func (q *queryStruct) DusunListName() []string {
	names := make([]string, 0)
	query := `
	select d."name" from dusun d order by d."name" asc
	`
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Rows(&names, sqlrows); err != nil {
		hlog.Info(err)
		return names

	}
	return names
}

func (q *queryStruct) RWListName() []string {
	names := make([]string, 0)
	query := `
	select distinct on (rw."name") rw."name" from rw order by rw."name" asc
	`
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Rows(&names, sqlrows); err != nil {
		hlog.Info(err)
		return names

	}
	return names
}

func (q *queryStruct) RTListName() []string {
	names := make([]string, 0)
	query := `
	select distinct on (rt."name") rt."name" from rt order by rt."name" asc 	
	`
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Rows(&names, sqlrows); err != nil {
		hlog.Info(err)
		return names

	}
	return names
}

func (q *queryStruct) GetUserByIdf(id int) (*datastruct.UtilGetUser, error) {
	data := new(datastruct.UtilGetUser)
	query := fmt.Sprintf(`
	select u."uuid", u.username, coalesce(j."name", '') as jabatan, coalesce(u.photo_url, '') as photo_url 
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	left join jabatan j on ud.jabatan_code = j.code
	where u.id = %v limit 1
	`, id)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Row(data, sqlrows); err != nil {
		hlog.Info(err)
		return data, err

	}
	return data, nil
}

func (q *queryStruct) GetUserById(id string) (*datastruct.UtilGetUser, error) {
	data := new(datastruct.UtilGetUser)
	query := fmt.Sprintf(`
	select u."uuid", u.username, coalesce(j."name", '') as jabatan, coalesce(u.photo_url, '') as photo_url 
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	left join jabatan j on ud.jabatan_code = j.code
	where u.uuid = '%v' limit 1
	`, id)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Row(data, sqlrows); err != nil {
		hlog.Info(err)
		return data, err

	}
	return data, nil
}

func (q *queryStruct) GetStatusPresensi(id string) (*string, error) {
	statusPresensi := new(string)

	query := fmt.Sprintf(`
	select p.status_presensi_code from presensi p 
	where p."uuid" = '%v' limit 1
	`, id)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Row(statusPresensi, sqlrows); err != nil {
		hlog.Info(err)
		return statusPresensi, err

	}
	return statusPresensi, nil
}

func (q *queryStruct) CheckPertanyaanSurvei(id string) (*string, error) {
	pertanyaanId := new(string)

	query := fmt.Sprintf(`
	select js.pertanyaan_uuid from jawaban_survei js 
	where js.pertanyaan_uuid = '%v' limit 1
	`, id)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	if err := scan.Row(pertanyaanId, sqlrows); err != nil {
		hlog.Info(err)
		return pertanyaanId, err

	}
	return pertanyaanId, nil
}

func (q *queryStruct) GetTotalSurveiByPertanyaanPilihan(id, pilihan string) (*float64, error) {
	total := new(float64)

	query := fmt.Sprintf(`
	select count(js."uuid") from jawaban_survei js 
	where js.pertanyaan_uuid = '%v' and js.jawaban = '%v' limit 1
	`, id, pilihan)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	scan.Row(total, sqlrows)

	totalF := RoundFloat(*total, 2)
	return &totalF, nil
}

func (q *queryStruct) GetTotalSurveiByPertanyaan(id string) (*float64, error) {
	total := new(float64)

	query := fmt.Sprintf(`
	select count(js."uuid") from jawaban_survei js 
	where js.pertanyaan_uuid = '%v' limit 1
	`, id)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	scan.Row(total, sqlrows)

	totalF := RoundFloat(*total, 2)
	return &totalF, nil
}

func (q *queryStruct) GetTotalSurveiByPertanyaanPilihanPeriod(id, pilihan string, mulai, akhir, tahun int) (*float64, error) {
	total := new(float64)

	query := fmt.Sprintf(`
	select count(js."uuid") from jawaban_survei js 
	where js.pertanyaan_uuid = '%v' and js.jawaban = '%v' 
	and extract(month from js.created_at) >= %v 
	and extract(month from js.created_at) <= %v
	and extract(year from js.created_at) = %v
	limit 1
	`, id, pilihan, mulai, akhir, tahun)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	scan.Row(total, sqlrows)

	totalF := RoundFloat(*total, 2)
	return &totalF, nil
}

func (q *queryStruct) GetTotalSurveiByPertanyaanPeriod(id string, mulai, akhir, tahun int) (*float64, error) {
	total := new(float64)

	query := fmt.Sprintf(`
	select count(js."uuid") from jawaban_survei js 
	where js.pertanyaan_uuid = '%v' 
	and extract(month from js.created_at) >= %v 
	and extract(month from js.created_at) <= %v
	and extract(year from js.created_at) = %v
	limit 1
	`, id, mulai, akhir, tahun)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		hlog.Info(err)
	}
	scan.Row(total, sqlrows)

	totalF := RoundFloat(*total, 2)
	return &totalF, nil
}
