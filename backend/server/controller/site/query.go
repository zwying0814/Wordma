package site

import (
	"github.com/gofiber/fiber/v2"
	"wordma/log"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleQuerySites(c *fiber.Ctx) error {
	var err error
	// 查询全部记录
	var sites []model.Site
	err = model.DB.Find(&sites).Error
	if err != nil {
		log.Fatalf("failed to query sites: %v", err)
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SendSuccess(c, sites)
}
