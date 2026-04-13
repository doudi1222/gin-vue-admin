package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/router"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(routerGroup *gin.RouterGroup) {
	// 注册课程路由
	router.RouterGroupApp.CourseRouter.InitCourseRouter(routerGroup)
	
	// 注册章节和课时路由
	router.RouterGroupApp.ChapterLessonRouter.InitChapterLessonRouter(routerGroup)
	
	// 注册报名路由
	router.RouterGroupApp.EnrollmentRouter.InitEnrollmentRouter(routerGroup)
	
	// 注册学习进度路由
	router.RouterGroupApp.ProgressRouter.InitProgressRouter(routerGroup)
	
	// 注册作业路由
	router.RouterGroupApp.HomeworkRouter.InitHomeworkRouter(routerGroup)
	
	// 注册证书路由
	router.RouterGroupApp.CertificateRouter.InitCertificateRouter(routerGroup)
	
	// 注册评论路由
	router.RouterGroupApp.CommentRouter.InitCommentRouter(routerGroup)
}
