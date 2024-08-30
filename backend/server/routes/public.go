package routes

import (
	"github.com/gofiber/fiber/v2"
	"wordma/server/controller/comment"
	"wordma/server/controller/post"
	"wordma/server/controller/user"
)

// RegisterPublicRoutes 注册公共路由
func RegisterPublicRoutes(api fiber.Router) {
	public := api.Group("/")

	// 用户登录
	public.Post("/login", user.HandleLogin)

	// 新增评论
	public.Post("/comment", comment.HandleCreateComment)
	// 获取评论列表
	public.Get("/comment", comment.HandleQueryComments)
	// 评论点踩
	public.Post("/comment/vote", comment.HandleCommentVote)

	// 文章点踩
	public.Post("/post/vote", post.HandlePostVote)
	// 文章阅读量统计
	public.Post("/post/view", post.HandlePostView)
	// 文章全部信息
	public.Get("/post", post.HandleQueryPost)

}
