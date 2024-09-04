package comment

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/utils"
)

func HandleCommentEmoji(c *fiber.Ctx) error {
	// 返回全部emoji数据到前端
	return utils.SendSuccess(c, utils.EmojiJson)
}
