package routes

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/controller/comment"
	"wordma/server/controller/site"
	"wordma/server/middleware"
)

func RegisterAdminRoutes(api fiber.Router) {
	auth := api.Group("/", middleware.AuthMiddleware)
	// 获取全部站点
	auth.Get("/sites", site.HandleQuerySites)
	// 新增站点
	auth.Post("/site", site.HandleCreateSite)
	// 删除站点
	auth.Delete("/site/:id", site.HandleDeleteSite)
	// 更新站点
	auth.Put("/site/:id", site.HandleUpdateSite)

	// 根据站点ID查询所有评论信息
	auth.Get("/site/comments", site.HandleGetSiteCommentsList)
	// 编辑评论信息
	auth.Put("/comment/:id", comment.HandleUpdateComments)
	// 删除评论信息
	auth.Delete("/comment/:id", comment.HandleDeleteComments)
}
