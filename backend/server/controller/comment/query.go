package comment

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/88250/lute"
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	"gorm.io/gorm"
	"strings"
	"wordma/config"
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
	if site == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "站点不存在")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}

	// 查询文章是否存在
	post, err := model.GetPostBySlug(data.PostSlug)
	if post == nil {
		return utils.SendError(c, fiber.StatusBadRequest, "文章不存在")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
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

	if data.PageSize > 0 {
		query = query.Limit(data.PageSize)
	} else {
		query = query.Limit(10)
	}
	if data.PageNumber > 0 {
		query = query.Offset((data.PageNumber - 1) * data.PageSize)
	} else {
		query = query.Offset(0)
	}

	if err := query.Find(&topLevelComments).Error; err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询出错"+err.Error())
	}

	// 递归加载子评论
	var commentDTOs []dto.ResponseCommentListDTO
	for _, comment := range topLevelComments {
		commentDTOs = append(commentDTOs, buildCommentDTO(comment, ""))
	}

	return utils.SendResponse(c, fiber.StatusOK, "查询成功", fiber.Map{
		"comments": commentDTOs,
		"total":    len(commentDTOs),
	})

}

// 递归函数用于构建 CommentDTO 并加载子评论
func buildCommentDTO(comment model.Comment, parentAuthorName string) dto.ResponseCommentListDTO {
	var replies []model.Comment
	model.DB.Preload("User").Preload("Post").Where("parent = ?", comment.ID).Find(&replies)

	var replyDTOs []dto.ResponseCommentListDTO
	for _, reply := range replies {
		replyDTOs = append(replyDTOs, buildCommentDTO(reply, comment.User.Name))
	}

	ua := useragent.Parse(comment.UA)

	if config.SupportMarkdown {
		luteEngine := lute.New()
		comment.Content = luteEngine.MarkdownStr("comment", comment.Content)
	}

	return dto.ResponseCommentListDTO{
		ID:           comment.ID,
		Content:      replaceFlagWithImagePath(comment.Content),
		Region:       comment.Region,
		OS:           ua.OS,
		Browser:      ua.Name,
		Type:         comment.Type,
		Up:           comment.Up,
		Down:         comment.Down,
		UserID:       comment.UserID,
		UserName:     comment.User.Name,
		UserAvatar:   getCravatarURL(comment.User.Email, 80),
		IsAdmin:      comment.User.Role == "admin",
		PostSlug:     comment.Post.Slug,
		Parent:       comment.Parent,
		ParentAuthor: parentAuthorName,
		CreatedAt:    comment.CreatedAt.Format("2006年1月2日 15:04"),
		Replies:      replyDTOs,
	}
}

// getCravatarURL 生成 Cravatar 头像的 URL
func getCravatarURL(email string, size int) string {
	// 将电子邮件地址转为小写并去除空格
	trimmedEmail := strings.TrimSpace(strings.ToLower(email))

	// 生成 MD5 散列
	hash := md5.Sum([]byte(trimmedEmail))
	hashStr := hex.EncodeToString(hash[:])

	// 构建 Cravatar URL
	cravatarURL := fmt.Sprintf("https://cravatar.cn/avatar/%s?s=%d", hashStr, size)

	return cravatarURL
}

// 解析Content中的emoji符号
func replaceFlagWithImagePath(input string) string {
	for _, emoji := range utils.EmojiJson {
		for _, image := range emoji.Images {
			if strings.Contains(input, image.Flag) {
				// 构建完整的 HTML <img> 标签路径
				imageTag := `<img emoji src="` + emoji.Path + image.Icon + `" alt="` + image.Text + `" />`
				// 替换 flag 为 HTML <img> 标签
				input = strings.ReplaceAll(input, image.Flag, imageTag)
			}
		}
	}
	return input
}
