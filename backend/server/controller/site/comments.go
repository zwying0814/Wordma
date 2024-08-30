package site

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleGetSiteCommentsList(c *fiber.Ctx) error {
	var err error
	var data dto.ReceiveCommentListBackendDTO

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

	// 查找与 SiteID 相关的所有 Post ID
	var postIDs []uint
	err = model.DB.Model(&model.Post{}).Where("site_id = ?", site.ID).Pluck("id", &postIDs).Error
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "无法获取文章列表: "+err.Error())
	}

	if len(postIDs) == 0 {
		return utils.SendResponse(c, fiber.StatusOK, "暂无文章", []model.Comment{})
	}

	// 构建查询条件
	query := model.DB.Model(&model.Comment{}).Where("post_id IN ?", postIDs)

	// 搜索条件
	if data.Search != "" {
		query = query.Where("content LIKE ?", "%"+data.Search+"%")
	}

	// 排序条件
	switch data.SortBy {
	case "date_asc":
		query = query.Order("created_at ASC")
	case "date_desc":
		query = query.Order("created_at DESC")
	case "vote":
		query = query.Order("up DESC, down ASC")
	default:
		query = query.Order("created_at DESC") // 默认按日期倒序排序
	}

	// 分页处理
	if data.Limit > 0 {
		query = query.Limit(data.Limit)
	} else {
		query = query.Limit(10)
	}
	if data.Offset > 0 {
		query = query.Offset(data.Offset)
	} else {
		query = query.Offset(0)
	}

	// 执行查询
	var comments []model.Comment
	err = query.Preload("Post").Preload("User").Find(&comments).Error
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "无法获取评论列表: "+err.Error())
	}

	// 返回结果
	return utils.SendResponse(c, fiber.StatusOK, "加载评论成功", comments)
}
