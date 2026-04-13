package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// EnrollmentRouter 报名路由
type EnrollmentRouter struct{}

// InitEnrollmentRouter 初始化报名路由
func (r *EnrollmentRouter) InitEnrollmentRouter(router *gin.RouterGroup) {
	enrollmentRouter := router.Group("/enrollment")
	{
		enrollmentRouter.POST("", api.ApiGroupApp.EnrollmentApi.EnrollCourse)
		enrollmentRouter.GET("/my-courses", api.ApiGroupApp.EnrollmentApi.GetUserCourses)
		enrollmentRouter.GET("/course-students/:courseId", api.ApiGroupApp.EnrollmentApi.GetCourseStudents)
		enrollmentRouter.POST("/cancel", api.ApiGroupApp.EnrollmentApi.CancelEnrollment)
	}
}
