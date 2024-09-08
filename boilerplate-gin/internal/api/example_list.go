package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary      List
// @Tags       	 Example
// @Produce      json
// @Param        order		query   string	false  "desc/asc default(desc)"
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /example/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleList(c *gin.Context) {
	ctx := context.Background()
	queryParam := new(dto.PaginationReq).Init()
	if err := c.BindQuery(&queryParam); err != nil {
		util.NewResponse(c).Error(err, "", 400)
		return
	}
	pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	data, countRow, err := m.exampleService.List(&ctx, pageMeta.Limit, pageMeta.Offset)
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  countRow,
		CountPages: util.NewPagination().GetCountPages(countRow, pageMeta.Limit),
	}
	util.NewResponse(c).Success(data, meta, infrastructure.Localize("OK_READ"))
}
