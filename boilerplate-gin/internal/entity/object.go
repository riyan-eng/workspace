package entity

type ServObjectCreate struct {
	Id          *string
	Name        *string
	Owner       *string
	Size        *int
	ContentType *string
	Url         *string
	Path        *string
}

type ServObjectDetail struct {
	Id *string
}

type ServObjectDelete struct {
	Id *string
}
