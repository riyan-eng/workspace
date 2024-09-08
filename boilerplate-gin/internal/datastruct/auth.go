package datastruct

import "github.com/golang-jwt/jwt/v5"

type AuthLoginData struct {
	Id          string `db:"uuid"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	RoleCode    string `db:"role_code"`
	RoleName    string `db:"role_name"`
	JabatanName any    `db:"jabatan_name"`
	IsActive    bool   `db:"is_active"`
}

type AuthToken struct {
	AccessToken    *string
	AccessExpired  *jwt.NumericDate
	RefreshToken   *string
	RefreshExpired *jwt.NumericDate
}

type AuthMe struct {
	UUID         string `db:"uuid" json:"id"`
	Username     string `db:"username" json:"username"`
	RoleCode     string `db:"role_code"`
	RoleName     string `db:"role_name"`
	JabatanCode  any    `db:"jabatan_code" json:"jabatan_code"`
	JabatanName  any    `db:"jabatan_name" json:"jabatan_name"`
	TempatLahir  any    `db:"birth_place" json:"tempat_lahir"`
	TanggalLahir any    `db:"birth_date" json:"tanggal_lahir"`
	Alamat       any    `db:"address" json:"alamat"`
	PhotoUrl     any    `db:"photo_url" json:"photo_url"`
}
