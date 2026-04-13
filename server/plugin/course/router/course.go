package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// CourseRouter 课程路由
type CourseRouter struct{}

// InitCourseRouter 初始化课程路由
func (r *CourseRouter) InitCourseRouter(router *gin.RouterGroup) {
	courseRouter := router.Group("/course")
	{
		courseRouter.GET("/list", api.ApiGroupApp.CourseApi.GetCourseList)
		courseRouter.GET("/:id", api.ApiGroupApp.CourseApi.GetCourseByID)
		courseRouter.POST("", api.ApiGroupApp.CourseApi.CreateCourse)
		courseRouter.PUT("", api.ApiGroupApp.CourseApi.UpdateCourse)
		courseRouter.DELETE("/:id", api.ApiGroupApp.CourseApi.DeleteCourse)
	}
}
