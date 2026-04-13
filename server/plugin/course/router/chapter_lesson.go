package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// ChapterLessonRouter 章节和课时路由
type ChapterLessonRouter struct{}

// InitChapterLessonRouter 初始化章节和课时路由
func (r *ChapterLessonRouter) InitChapterLessonRouter(router *gin.RouterGroup) {
	// 章节路由
	chapterRouter := router.Group("/chapter")
	{
		chapterRouter.POST("", api.ApiGroupApp.ChapterLessonApi.CreateChapter)
		chapterRouter.PUT("", api.ApiGroupApp.ChapterLessonApi.UpdateChapter)
		chapterRouter.DELETE("/:id", api.ApiGroupApp.ChapterLessonApi.DeleteChapter)
	}

	// 课时路由
	lessonRouter := router.Group("/lesson")
	{
		lessonRouter.POST("", api.ApiGroupApp.ChapterLessonApi.CreateLesson)
		lessonRouter.PUT("", api.ApiGroupApp.ChapterLessonApi.UpdateLesson)
		lessonRouter.DELETE("/:id", api.ApiGroupApp.ChapterLessonApi.DeleteLesson)
	}
}
