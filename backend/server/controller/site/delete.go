package site

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleDeleteSite(c *fiber.Ctx) error {
	var err error
	id, _ := c.ParamsInt("id")
	site, err := model.FindSiteByID(uint(id))
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	err = model.DeleteSiteById(uint(id))
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	return utils.SendSuccess(c, "删除成功")
}
