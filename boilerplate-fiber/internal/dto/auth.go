package dto

type AuthLogin struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type AuthRefresh struct {
	Token string `json:"token" valid:"required"`
}
