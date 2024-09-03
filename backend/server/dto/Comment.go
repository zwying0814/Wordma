package dto

// CommentDTO 接收前端评论参数
type CommentDTO struct {
	Name     string `json:"name" validate:"required"`      // 昵称
	Email    string `json:"email" validate:"required"`     // 邮箱
	Url      string `json:"url" validate:"optional"`       // 链接
	Content  string `json:"content" validate:"required"`   // 评论内容
	Parent   uint   `json:"parent" validate:"optional"`    // 父级ID
	UA       string `json:"ua" validate:"optional"`        // 浏览器UA
	PostSlug string `json:"post_slug" validate:"required"` // 文章Slug
	SiteID   uint   `json:"site_id" validate:"required"`   // 站点ID
}

// UpdateCommentDTO 接收修改评论的参数，目前只允许修改内容
type UpdateCommentDTO struct {
	Content string `json:"content" validate:"optional"` // 评论内容
	Type    string `json:"type" validate:"optional"`    // 评论类型
}

// ReceiveCommentListDTO 加载评论接收参数
type ReceiveCommentListDTO struct {
	PostSlug string `query:"post_slug" json:"post_slug" validate:"required"` // 文章Slug
	SiteID   uint   `query:"site_id" json:"site_id" validate:"required"`     // 站点ID

	PageNumber int    `query:"page_number" json:"page_number" validate:"optional"`
	PageSize   int    `query:"page_size" json:"page_size" validate:"optional"`
	SortBy     string `query:"sort_by" json:"sort_by" enums:"date_asc,date_desc,vote" validate:"optional"` // Sort by condition
	Search     string `query:"search" json:"search" validate:"optional"`                                   // Search keywords
}

// ReceiveCommentListBackendDTO 后端加载评论接收参数
type ReceiveCommentListBackendDTO struct {
	SiteID uint `query:"site_id" json:"site_id" validate:"optional"` // 站点ID

	Limit  int    `query:"limit" json:"limit" validate:"optional"`                                     // The limit for pagination
	Offset int    `query:"offset" json:"offset" validate:"optional"`                                   // The offset for pagination
	SortBy string `query:"sort_by" json:"sort_by" enums:"date_asc,date_desc,vote" validate:"optional"` // Sort by condition
	Search string `query:"search" json:"search" validate:"optional"`
}

type ResponseCommentListDTO struct {
	ID           uint                     `json:"id"`
	Content      string                   `json:"content"`
	OS           string                   `json:"os"`
	Browser      string                   `json:"browser"`
	Region       string                   `json:"region"`
	Type         string                   `json:"type"`
	Up           int                      `json:"up"`
	Down         int                      `json:"down"`
	UserID       uint                     `json:"user_id"`
	UserName     string                   `json:"user_name"`
	UserAvatar   string                   `json:"user_avatar"`
	IsAdmin      bool                     `json:"is_admin"`
	PostSlug     string                   `json:"post_slug"`
	Parent       uint                     `json:"parent"`
	ParentAuthor string                   `json:"parent_author"`
	CreatedAt    string                   `json:"created_at"`
	Replies      []ResponseCommentListDTO `json:"replies,omitempty"` // 子评论的递归结构
}

type CommentVoteDTO struct {
	CommentID uint   `json:"comment_id" validate:"required"`
	VoteType  string `json:"vote_type" validate:"required,oneof=up down"`
}
