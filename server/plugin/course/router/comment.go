package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// CommentRouter 评论路由
type CommentRouter struct{}

// InitCommentRouter 初始化评论路由
func (r *CommentRouter) InitCommentRouter(router *gin.RouterGroup) {
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("", api.ApiGroupApp.CommentApi.CreateComment)
		commentRouter.GET("/course/:courseId", api.ApiGroupApp.CommentApi.GetCourseComments)
		commentRouter.DELETE("/:id", api.ApiGroupApp.CommentApi.DeleteComment)
	}
}
