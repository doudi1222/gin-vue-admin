package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// ChapterLessonApi 章节和课时API
type ChapterLessonApi struct{}

// CreateChapter 创建章节
// @Tags     Chapter
// @Summary  创建章节
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.CreateChapterRequest true "创建章节参数"
// @Success  200  {object} response.Response{data=model.Chapter} "创建成功"
// @Router   /chapter [post]
func (a *ChapterLessonApi) CreateChapter(c *gin.Context) {
	var req request.CreateChapterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chapter, err := service.ServiceGroupApp.ChapterLessonService.CreateChapter(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(chapter, "创建成功", c)
}

// UpdateChapter 更新章节
// @Tags     Chapter
// @Summary  更新章节
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UpdateChapterRequest true "更新章节参数"
// @Success  200  {object} response.Response{data=model.Chapter} "更新成功"
// @Router   /chapter [put]
func (a *ChapterLessonApi) UpdateChapter(c *gin.Context) {
	var req request.UpdateChapterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chapter, err := service.ServiceGroupApp.ChapterLessonService.UpdateChapter(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(chapter, "更新成功", c)
}

// DeleteChapter 删除章节
// @Tags     Chapter
// @Summary  删除章节
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "章节ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /chapter/{id} [delete]
func (a *ChapterLessonApi) DeleteChapter(c *gin.Context) {
	id := c.Param("id")
	chapterID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的章节ID", c)
		return
	}

	err = service.ServiceGroupApp.ChapterLessonService.DeleteChapter(uint(chapterID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// CreateLesson 创建课时
// @Tags     Lesson
// @Summary  创建课时
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.CreateLessonRequest true "创建课时参数"
// @Success  200  {object} response.Response{data=model.Lesson} "创建成功"
// @Router   /lesson [post]
func (a *ChapterLessonApi) CreateLesson(c *gin.Context) {
	var req request.CreateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	lesson, err := service.ServiceGroupApp.ChapterLessonService.CreateLesson(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(lesson, "创建成功", c)
}

// UpdateLesson 更新课时
// @Tags     Lesson
// @Summary  更新课时
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UpdateLessonRequest true "更新课时参数"
// @Success  200  {object} response.Response{data=model.Lesson} "更新成功"
// @Router   /lesson [put]
func (a *ChapterLessonApi) UpdateLesson(c *gin.Context) {
	var req request.UpdateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	lesson, err := service.ServiceGroupApp.ChapterLessonService.UpdateLesson(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(lesson, "更新成功", c)
}

// DeleteLesson 删除课时
// @Tags     Lesson
// @Summary  删除课时
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "课时ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /lesson/{id} [delete]
func (a *ChapterLessonApi) DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	lessonID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课时ID", c)
		return
	}

	err = service.ServiceGroupApp.ChapterLessonService.DeleteLesson(uint(lessonID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
