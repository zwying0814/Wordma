package post

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandlePostVote(c *fiber.Ctx) error {
	var data dto.PostVoteDTO

	// 解析并验证传入参数
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}
	// 查找 Site
	site, err := model.FindSiteByID(data.SiteID)
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}

	// 查找文章
	post, err := model.GetPostBySlug(data.PostSlug)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理非未找到记录的错误
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询错误: "+err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果文章未找到，则创建新文章
		post = &model.Post{
			Slug:   data.PostSlug,
			SiteID: data.SiteID,
			Up:     0,
			Down:   0,
		}
		if data.VoteType == "up" {
			post.Up = 1
		} else if data.VoteType == "down" {
			post.Down = 1
		}

		// 保存新创建的文章
		err = model.CreatePost(post)
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, "无法创建新文章: "+err.Error())
		}
	} else {
		// 如果找到文章，根据投票类型更新 Up 或 Down 计数
		if data.VoteType == "up" {
			post.Up += 1
		} else if data.VoteType == "down" {
			post.Down += 1
		}

		// 保存更新后的文章
		if err := model.DB.Save(post).Error; err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, "无法保存投票结果: "+err.Error())
		}
	}

	// 返回成功响应
	return utils.SendResponse(c, fiber.StatusOK, "OK", fiber.Map{
		"post_id": post.ID,
		"up":      post.Up,
		"down":    post.Down,
	})
}
