package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// HomeworkRouter 作业路由
type HomeworkRouter struct{}

// InitHomeworkRouter 初始化作业路由
func (r *HomeworkRouter) InitHomeworkRouter(router *gin.RouterGroup) {
	homeworkRouter := router.Group("/homework")
	{
		homeworkRouter.POST("", api.ApiGroupApp.HomeworkApi.CreateHomework)
		homeworkRouter.PUT("", api.ApiGroupApp.HomeworkApi.UpdateHomework)
		homeworkRouter.DELETE("/:id", api.ApiGroupApp.HomeworkApi.DeleteHomework)
		homeworkRouter.GET("/course/:courseId", api.ApiGroupApp.HomeworkApi.GetCourseHomeworks)
		homeworkRouter.POST("/submit", api.ApiGroupApp.HomeworkApi.SubmitHomework)
		homeworkRouter.PUT("/grade", api.ApiGroupApp.HomeworkApi.GradeHomework)
		homeworkRouter.GET("/submissions/:homeworkId", api.ApiGroupApp.HomeworkApi.GetHomeworkSubmissions)
		homeworkRouter.GET("/my-submission/:homeworkId", api.ApiGroupApp.HomeworkApi.GetUserHomeworkSubmission)
	}
}
