package comment

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleDeleteComments(c *fiber.Ctx) error {
	var err error
	id, _ := c.ParamsInt("id")
	comment, err := model.GetCommentByID(uint(id))
	if comment == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "评论不存在")
	}
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	err = model.DeleteCommentByID(uint(id))
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	return utils.SendSuccess(c, "删除成功")
}
