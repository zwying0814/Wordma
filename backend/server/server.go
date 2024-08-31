package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"wordma/config"
	"wordma/server/model"
	"wordma/server/routes"
	"wordma/server/utils"
)

// Bootstrap 启动后端服务
func Bootstrap() (*fiber.App, error) {

	// 初始化Fiber
	fb := fiber.New(fiber.Config{
		Immutable:          true,
		BodyLimit:          3 * 1024 * 1024, // MB
		StreamRequestBody:  true,
		EnableIPValidation: true,
	})

	// 初始化路由
	routes.InitRoutes(fb)

	if config.NeedFilter {
		// 初始化过滤器
		utils.InitCommentFilter()
	}

	if config.DevelopMode {
		// 开发模式，插入一条测试数据
		err := model.CreateSite(&model.Site{
			Url:  "test.com",
			Name: "测试",
		})
		if err != nil {
			panic("插入测试数据错误" + err.Error())
		}
	}

	// 配置跨域
	fb.Use(cors.New())

	// 监听端口
	return fb, fb.Listen(":" + config.Port)
}
