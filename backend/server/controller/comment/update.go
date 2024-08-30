package comment

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleUpdateComments(c *fiber.Ctx) error {
	var err error
	id, _ := c.ParamsInt("id")
	var data dto.UpdateCommentDTO
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}
	comment, err := model.GetCommentByID(uint(id))
	if comment == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "评论不存在")
	}
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	commentData := &model.Comment{
		Content: data.Content,
		Type:    data.Type,
	}
	err = model.UpdateCommentByID(uint(id), commentData)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	return utils.SendSuccess(c, "更新成功")
}
