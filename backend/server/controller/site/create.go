package site

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

// HandleCreateSite 新增一个站点
func HandleCreateSite(c *fiber.Ctx) error {
	var data dto.ReceiveCreateSiteDTO
	var err error
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}
	site, err := model.FindSiteByURL(data.Url)
	if site != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点已存在")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询错误")
	}
	newSite := model.Site{
		Name: data.Name,
		Url:  data.Url,
	}
	if err := model.CreateSite(&newSite); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库插入错误")
	}
	return utils.SendResponse(c, fiber.StatusOK, "站点创建成功", newSite)
}
