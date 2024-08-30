package comment

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleCommentVote(c *fiber.Ctx) error {

	var data dto.CommentVoteDTO

	// 解析并验证传入参数
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}

	// 查找评论
	var comment model.Comment
	if err := model.DB.First(&comment, data.CommentID).Error; err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "评论不存在")
	}

	// 根据投票类型更新评论的 Up 或 Down 计数
	if data.VoteType == "up" {
		comment.Up += 1
	} else if data.VoteType == "down" {
		comment.Down += 1
	}

	// 保存更新后的评论
	if err := model.DB.Save(&comment).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}

	// 返回成功响应
	return utils.SendResponse(c, fiber.StatusOK, "投票成功", fiber.Map{
		"comment_id": comment.ID,
		"up":         comment.Up,
		"down":       comment.Down,
	})
}
