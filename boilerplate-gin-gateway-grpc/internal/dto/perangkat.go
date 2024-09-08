package dto

type PerangkatCreate struct {
	Username    string `json:"username" valid:"required"`
	JabatanCode string `json:"jabatan_code" valid:"required;in:A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q"`
	BirthPlace  string `json:"tempat_lahir" valid:"required"`
	BirthDate   string `json:"tanggal_lahir" valid:"required;date:yyyy-mm-dd"`
	Address     string `json:"alamat" valid:"required"`
	PhotoUrl    string `json:"photo_url" valid:"required"`
}

type PerangkatPatch struct {
	Username    string `json:"username" valid:"required"`
	JabatanCode string `json:"jabatan_code" valid:"required;in:A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q"`
	BirthPlace  string `json:"tempat_lahir" valid:"required"`
	BirthDate   string `json:"tanggal_lahir" valid:"required;date:yyyy-mm-dd"`
	Address     string `json:"alamat" valid:"required"`
	PhotoUrl    string `json:"photo_url" valid:"required"`
}

type PerangkatResetPassword struct {
	Password string `json:"password" valid:"required;min:8"`
}
