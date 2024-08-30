package dto

type PostVoteDTO struct {
	SiteID   uint   `query:"site_id" json:"site_id" validate:"required"` // 站点ID
	PostSlug string `json:"post_slug" validate:"required"`
	VoteType string `json:"vote_type" validate:"required,oneof=up down"`
}

type PostViewDTO struct {
	SiteID   uint   `query:"site_id" json:"site_id" validate:"required"` // 站点ID
	PostSlug string `json:"post_slug" validate:"required"`
}

// PostListDTO 站点下文章全部数据
type PostListDTO struct {
	SiteID uint `query:"site_id" json:"site_id" validate:"required"` // 站点ID

	Limit  int `query:"limit" json:"limit" validate:"optional"`   // The limit for pagination
	Offset int `query:"offset" json:"offset" validate:"optional"` // The offset for pagination
}
