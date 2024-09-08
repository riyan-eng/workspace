package dto

type PaginationReq struct {
	Page   int    `form:"page"`
	Limit  int    `form:"per_page"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

func (p PaginationReq) Init() PaginationReq {
	p.Page = 1
	p.Limit = 10
	p.Order = "desc"
	return p
}
