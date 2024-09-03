package dto

type PostVoteDTO struct {
	SiteID   uint   `query:"site_id" json:"site_id" validate:"required"` // 站点ID
	PostSlug string `json:"post_slug" validate:"required"`
	VoteType string `json:"vote_type" validate:"required,oneof=up down"`
}

type PostDataDTO struct {
	SiteID   uint   `json:"site_id" validate:"required"` // 站点ID
	PostSlug string `json:"post_slug" validate:"required"`
}

// PostListDTO 站点下文章全部数据
type PostListDTO struct {
	SiteID uint `query:"site_id" json:"site_id" validate:"required"` // 站点ID

	PageNumber int `query:"page_number" json:"page_number" validate:"optional"`
	PageSize   int `query:"page_size" json:"page_size" validate:"optional"`
}
