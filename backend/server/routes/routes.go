package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")

	// v1版本
	v1 := api.Group("/v1")

	// 注册公共路由
	RegisterPublicRoutes(v1)
	// 注册后台路由
	RegisterAdminRoutes(v1)

	fmt.Println("路由初始化成功！")

}
