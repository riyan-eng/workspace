package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Upload
// @Tags        Object
// @Param   	file  formData     file     true  "File"
// @Router		/object/ [post]
// @Security ApiKeyAuth
func (m *ServiceServer) ObjectUpload(c *fiber.Ctx) error {
	ctx := context.Background()
	user := util.CurrentUser(c)
	file, errT := c.FormFile("file")
	if errT != nil {
		return util.NewResponse(c).Error(errT.Error(), "", 400)

	}

	dataFile, err := util.NewFile(c).SaveLocal(file)
	if err.Errors != nil {
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

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
		return util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)

	}
	data := map[string]any{
		"id":  dataFile.Id,
		"url": dataFile.Url,
	}

	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"))
}
