package dto

// ReceiveCreateSiteDTO 前端创建站点数据
type ReceiveCreateSiteDTO struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}
