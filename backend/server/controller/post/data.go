package post

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleQueryPost(c *fiber.Ctx) error {
	var err error
	var data dto.PostListDTO

	// 解析并验证传入参数
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}

	// 查找 Site，确认站点是否存在
	site, err := model.FindSiteByID(data.SiteID)
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错: "+err.Error())
	}

	// 构建查询条件，从 Post 表中根据 SiteID 查询所需字段
	var posts []struct {
		Slug string `json:"slug"`
		Up   int    `json:"up"`
		Down int    `json:"down"`
		Read int    `json:"read"`
	}

	query := model.DB.Model(&model.Post{}).Where("site_id = ?", data.SiteID).Select("slug, up, down, read")

	// 分页处理
	if data.Limit > 0 {
		query = query.Limit(data.Limit)
	} else {
		query = query.Limit(10)
	}
	if data.Offset > 0 {
		query = query.Offset(data.Offset)
	}

	// 执行查询
	err = query.Scan(&posts).Error
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "无法获取文章列表: "+err.Error())
	}

	// 如果没有文章
	if len(posts) == 0 {
		return utils.SendResponse(c, fiber.StatusOK, "暂无文章", []model.Post{})
	}

	// 返回查询结果
	return utils.SendResponse(c, fiber.StatusOK, "获取文章成功", posts)
}
