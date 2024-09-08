package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary      List
// @Tags       	 Perangkat
// @Produce      json
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /perangkat/ [get]
// @Security ApiKeyAuth
func (m *ServiceServer) PerangkatList(c *fiber.Ctx) error {
	ctx := context.Background()
	queryParam := new(dto.PaginationReq).Init()
	if err := c.QueryParser(&queryParam); err != nil {
		return util.NewResponse(c).Error(err, "", 400)

	}
	pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	data, countRow, err := m.perangkatService.List(&ctx, &entity.ServPerangkatList{
		Search: &queryParam.Search,
		Limit:  pageMeta.Limit,
		Offset: pageMeta.Offset,
	})
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  countRow,
		CountPages: util.NewPagination().GetCountPages(countRow, pageMeta.Limit),
	}
	return util.NewResponse(c).Success(data, meta, infrastructure.Localize("OK_READ"))
}
