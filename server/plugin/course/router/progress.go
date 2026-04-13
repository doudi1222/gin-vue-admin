package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// ProgressRouter 学习进度路由
type ProgressRouter struct{}

// InitProgressRouter 初始化学习进度路由
func (r *ProgressRouter) InitProgressRouter(router *gin.RouterGroup) {
	progressRouter := router.Group("/progress")
	{
		progressRouter.PUT("", api.ApiGroupApp.ProgressApi.UpdateProgress)
		progressRouter.GET("/course/:courseId", api.ApiGroupApp.ProgressApi.GetCourseProgress)
		progressRouter.GET("/lesson/:lessonId", api.ApiGroupApp.ProgressApi.GetLessonProgress)
	}
}
