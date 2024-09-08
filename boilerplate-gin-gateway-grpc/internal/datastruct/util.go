package datastruct

type UtilGetUser struct {
	Id       string `db:"uuid"`
	Username string `db:"username"`
	Jabatan  string `db:"jabatan"`
	PhotoUrl string `db:"photo_url"`
}
