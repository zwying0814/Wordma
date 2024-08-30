package comment

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleQueryComments(c *fiber.Ctx) error {
	var err error
	var data dto.ReceiveCommentListDTO

	// 解析并验证传入参数
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}

	// 查询站点是否存在
	site, err := model.FindSiteByID(data.SiteID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}

	// 查询文章是否存在
	post, err := model.GetPostBySlug(data.PostSlug)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}
	if post == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "文章不存在")
	}

	// 如果站点文章都有，还需要看看站点下有没有这篇文章，没有也要返回错误
	if post.SiteID != site.ID {
		return utils.SendError(c, fiber.StatusBadRequest, "文章不属于该站点")
	}

	// 查询顶级评论
	var topLevelComments []model.Comment
	query := model.DB.Preload("User").Preload("Post").Where("post_id = ? AND parent = 0", post.ID)

	// 看看是不是管理员
	isAdmin := utils.CheckIsAdminReq(c)
	// 不是管理员只能看到type=published的评论
	if !isAdmin {
		query = query.Where("type = 'published'")
	}

	if data.Search != "" {
		query = query.Where("content LIKE ?", "%"+data.Search+"%")
	}

	switch data.SortBy {
	case "date_asc":
		query = query.Order("created_at ASC")
	case "date_desc":
		query = query.Order("created_at DESC")
	case "vote":
		query = query.Order("up - down DESC")
	default:
		query = query.Order("created_at DESC")
	}

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

	if err := query.Find(&topLevelComments).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}

	// 递归加载子评论
	var commentDTOs []dto.ResponseCommentListDTO
	for _, comment := range topLevelComments {
		commentDTOs = append(commentDTOs, buildCommentDTO(comment))
	}

	return utils.SendResponse(c, fiber.StatusOK, "查询成功", fiber.Map{
		"comments": commentDTOs,
		"total":    len(commentDTOs),
	})

}

// 递归函数用于构建 CommentDTO 并加载子评论
func buildCommentDTO(comment model.Comment) dto.ResponseCommentListDTO {
	var replies []model.Comment
	model.DB.Where("parent = ?", comment.ID).Find(&replies)

	var replyDTOs []dto.ResponseCommentListDTO
	for _, reply := range replies {
		replyDTOs = append(replyDTOs, buildCommentDTO(reply))
	}

	return dto.ResponseCommentListDTO{
		ID:        comment.ID,
		Content:   comment.Content,
		UA:        comment.UA,
		IP:        comment.IP,
		Region:    comment.Region,
		Type:      comment.Type,
		Up:        comment.Up,
		Down:      comment.Down,
		UserID:    comment.UserID,
		UserName:  comment.User.Name,
		UserEmail: comment.User.Email,
		PostSlug:  comment.Post.Slug,
		Parent:    comment.Parent,
		CreatedAt: comment.CreatedAt,
		Replies:   replyDTOs,
	}
}
