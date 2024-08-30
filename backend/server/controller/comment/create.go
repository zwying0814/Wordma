package comment

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"wordma/config"
	"wordma/log"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
	"wordma/server/utils/ip_region"
)

// HandleCreateComment 新增一条评论
func HandleCreateComment(c *fiber.Ctx) error {
	var err error
	var data dto.CommentDTO
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}
	//if strings.TrimSpace(data.Name) == "" {
	//	return utils.RespError(c, 400, "昵称不能为空", nil)
	//}
	//if strings.TrimSpace(data.Email) == "" {
	//	return utils.RespError(c, 400, "邮箱不能为空", nil)
	//}
	//if !utils.ValidateEmail(data.Email) {
	//	return utils.RespError(c, 400, "邮箱格式不正确", nil)
	//}
	//if data.Url != "" && !utils.ValidateURL(data.Url) {
	//	return utils.RespError(c, 400, "链接格式不正确", nil)
	//}
	site, err := model.FindSiteByID(data.SiteID)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "数据库查询错误"+err.Error())
	}
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}
	var (
		ip      = c.IP()
		ua      = string(c.Request().Header.UserAgent())
		isAdmin = utils.CheckIsAdminReq(c)
	)
	// 允许传入修正后的 UA
	if data.UA != "" {
		ua = data.UA
	}
	// 找到是那篇文章的，注意，如果找不到，那就需要插入新的文章记录，找不到的原因是第一条评论
	var post *model.Post
	post, err = model.FindOrCreatePost(data)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "数据库操作错误")
	}
	// 如果站点文章都有，还需要看看站点下有没有这篇文章，没有也要返回错误
	if post.SiteID != site.ID {
		return utils.SendError(c, fiber.StatusBadRequest, "文章不属于该站点")
	}
	// 检查评论是否合法
	if data.Parent != 0 {
		commentByID, err := model.GetCommentByID(data.Parent)
		if commentByID == nil {
			return utils.SendError(c, fiber.StatusBadRequest, "父级评论不存在")
		}
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, "数据库操作错误"+err.Error())
		}
		if commentByID.PostID != post.ID {
			return utils.SendError(c, fiber.StatusBadRequest, "父级评论不属于该文章")
		}
		if commentByID.Type == "noReply" {
			return utils.SendError(c, fiber.StatusBadRequest, "父级评论不允许回复")
		}
	}

	// 检查这个用户是否存在，不存在需要创建，通过邮箱查询，email作为唯一的标识，用户名可以不同
	user, err := utils.GetUserByReq(c)
	if errors.Is(err, utils.ErrTokenNotProvided) {
		user, err = model.FindOrCreateUser(data)
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, "数据库操作错误"+err.Error())
		}
		if user.Role == "ban" {
			return utils.SendError(c, fiber.StatusBadRequest, "该用户已被封禁")
		}
		// Update user
		err := model.UpdateUser(user.ID, fiber.Map{
			"url":     data.Url,
			"last_ip": ip,
			"last_ua": ua,
			"name":    data.Name,
			"email":   data.Email,
			"role":    "visitor",
		})
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, "数据库操作错误"+err.Error())
		}
		user, err = model.GetUserByID(user.ID)
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, "数据库操作错误"+err.Error())
		}
	} else if err != nil {
		// Login user error
		log.Error("[CommentCreate] Get user error: ", err)
		return utils.SendError(c, fiber.StatusBadRequest, "评论失败！"+err.Error())
	}
	comment := model.Comment{
		Content: data.Content,
		UA:      ua,
		IP:      ip,
		Region:  ip_region.IP2Region(ip),
		Type:    "published",
		Up:      0,
		Down:    0,
		PostID:  post.ID,
		UserID:  user.ID,
		Parent:  data.Parent,
	}
	if !isAdmin && config.NeedModeration {
		// 不是管理员评论 && 配置开启评论默认待审
		comment.Type = "pending"
	}

	// save to database
	if err := model.CreateComment(&comment); err != nil {
		log.Error("Save Comment error: ", err)
		return utils.SendError(c, fiber.StatusBadRequest, "评论失败！"+err.Error())
	}

	// todo:异步执行通知发送

	// 返回成功
	return utils.SendSuccess(c, "创建评论成功")
}
