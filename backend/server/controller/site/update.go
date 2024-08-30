package site

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleUpdateSite(c *fiber.Ctx) error {
	var err error
	id, _ := c.ParamsInt("id")
	var data dto.ReceiveCreateSiteDTO
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}
	site, err := model.FindSiteByID(uint(id))
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	site, err = model.FindSiteByURL(data.Url)
	if site != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Url已存在，请更改")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询错误")
	}
	siteData := &model.Site{
		Name: data.Name,
		Url:  data.Url,
	}
	err = model.UpdateSiteById(uint(id), siteData)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询错误")
	}
	return utils.SendSuccess(c, fiber.StatusOK)
}
