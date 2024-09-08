package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Upload
// @Tags        Object
// @Param   	file  formData     file     true  "File"
// @Router		/object/ [post]
// @Security ApiKeyAuth
func (m *ServiceServer) ObjectUpload(c *gin.Context) {
	ctx := context.Background()
	user := util.CurrentUser(c)
	file, errT := c.FormFile("file")
	if errT != nil {
		util.NewResponse(c).Error(errT.Error(), "", 400)
		return
	}

	dataFile, err := util.NewFile(c).SaveLocal(file)
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	if err := m.objectService.Create(&ctx, &entity.ServObjectCreate{
		Id:          dataFile.Id,
		Owner:       &user.UserId,
		Name:        dataFile.Name,
		Size:        dataFile.Size,
		ContentType: dataFile.ContentType,
		Url:         dataFile.Url,
		Path:        dataFile.Path,
	}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}
	data := map[string]any{
		"id":  dataFile.Id,
		"url": dataFile.Url,
	}

	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"))
}
