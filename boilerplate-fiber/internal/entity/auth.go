package entity

type ServAuthRegister struct {
	UserId   *string
	Email    *string
	UserName *string
	Password *string
	RoleCode *string
}

type ServAuthLogin struct {
	Username *string
	Password *string
}

type ServAuthRefresh struct {
	Token *string
}

type ServAuthResetPassword struct {
	Token    *string
	Password *string
}

type ServAuthResetPasswordToken struct {
	Email *string
}

type ServAuthResetPasswordTokenValidate struct {
	Token *string
}

type ServAuthLogout struct {
	UserId *string
}

type ServAuthMe struct {
	UserId *string
}
