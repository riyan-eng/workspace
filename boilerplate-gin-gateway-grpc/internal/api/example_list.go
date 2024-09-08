package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary      List
// @Tags       	 Example
// @Produce      json
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /example [get]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleList(c *gin.Context) {
	ctx := context.Background()
	queryParam := new(dto.PaginationReq).Init()
	if err := c.BindQuery(&queryParam); err != nil {
		util.NewResponse(c).Error(err, "", 400)
		return
	}
	pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	res, err := m.exampleRpcServer.List(ctx, &pb.TaskListRequest{
		Search: queryParam.Search,
		Limit:  int32(*pageMeta.Limit),
		Offset: int32(*pageMeta.Offset),	
	})

	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	var countRow int = int(res.TotalRows)

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  &countRow,
		CountPages: util.NewPagination().GetCountPages(&countRow, pageMeta.Limit),
	}
	util.NewResponse(c).Success(res.Data, meta, infrastructure.Localize("OK_READ"))
}
