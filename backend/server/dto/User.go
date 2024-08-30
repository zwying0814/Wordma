package dto

// ReceiveUserLoginDTO 前端登录数据
type ReceiveUserLoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
