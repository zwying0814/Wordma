package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"wordma/config"
	"wordma/server/routes"
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

	// 配置跨域
	fb.Use(cors.New())

	// 监听端口
	return fb, fb.Listen(":" + config.Port)
}
