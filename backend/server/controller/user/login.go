package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"wordma/config"
	"wordma/server/dto"
	"wordma/server/model"
	"wordma/server/utils"
)

func HandleLogin(c *fiber.Ctx) error {

	var data dto.ReceiveUserLoginDTO
	if isOK, resp := utils.ParamsDecode(c, &data); !isOK {
		return resp
	}
	// 查询用户
	user, err := model.GetUserByNameOrEmail(data.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.SendError(c, fiber.StatusBadRequest, "用户不存在")
	} else {
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, "数据库查询错误")
		}
	}
	// 密码比较
	if !user.CheckPassword(data.Password) {
		return utils.SendError(c, fiber.StatusBadRequest, "用户密码错误")
	}

	jwtToken, err := utils.LoginGetUserToken(*user, config.AppKey, 7*24*60*60)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}
	user.Password = ""
	return utils.SendSuccess(c, map[string]interface{}{
		"token": jwtToken,
		"user":  user,
	})
}
