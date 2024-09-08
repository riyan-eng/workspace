package datastruct

type PerangkatList struct {
	UUID         string `db:"uuid" json:"id"`
	Id           int    `db:"id" json:"-"`
	Idf          string `json:"idf"`
	Username     string `db:"username" json:"username"`
	JabatanName  any    `db:"jabatan_name" json:"jabatan_name"`
	TempatLahir  any    `db:"birth_place" json:"tempat_lahir"`
	TanggalLahir any    `db:"birth_date" json:"tanggal_lahir"`
	Alamat       any    `db:"address" json:"alamat"`
	PhotoUrl     any    `db:"photo_url" json:"photo_url"`
	TotalRows    int    `db:"total_rows" json:"-"`
}

type PerangkatDetail struct {
	UUID         string `db:"uuid" json:"id"`
	Id           int    `db:"id" json:"-"`
	Idf          string `json:"idf"`
	Username     string `db:"username" json:"username"`
	JabatanCode  any    `db:"jabatan_code" json:"jabatan_code"`
	JabatanName  any    `db:"jabatan_name" json:"jabatan_name"`
	TempatLahir  any    `db:"birth_place" json:"tempat_lahir"`
	TanggalLahir any    `db:"birth_date" json:"tanggal_lahir"`
	Alamat       any    `db:"address" json:"alamat"`
	PhotoUrl     any    `db:"photo_url" json:"photo_url"`
}
