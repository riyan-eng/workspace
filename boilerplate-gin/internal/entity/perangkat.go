package entity

type ServPerangkatList struct {
	Search *string
	Limit  *int
	Offset *int
}

type ServPerangkatCreate struct {
	Id          *string
	Username    *string
	Password    *string
	JabatanCode *string
	RoleCode    *string
	BirthPlace  *string
	BirthDate   *string
	Address     *string
	PhotoUrl    *string
}

type ServPerangkatPatch struct {
	Id          *string
	Username    *string
	RoleCode    *string
	JabatanCode *string
	BirthPlace  *string
	BirthDate   *string
	Address     *string
	PhotoUrl    *string
}

type ServPerangkatDetail struct {
	Id *string
}

type ServPerangkatDelete struct {
	Id *string
}

type ServPerangkatResetPassword struct {
	Id       *string
	Password *string
}
