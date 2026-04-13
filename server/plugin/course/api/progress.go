package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// ProgressApi 学习进度API
type ProgressApi struct{}

// UpdateProgress 更新学习进度
// @Tags     Progress
// @Summary  更新学习进度
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UpdateProgressRequest true "更新学习进度参数"
// @Success  200  {object} response.Response{data=model.LearningProgress} "更新成功"
// @Router   /progress [put]
func (a *ProgressApi) UpdateProgress(c *gin.Context) {
	var req request.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取
	req.UserID = userID

	progress, err := service.ServiceGroupApp.ProgressService.UpdateProgress(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(progress, "更新成功", c)
}

// GetCourseProgress 获取课程的学习进度
// @Tags     Progress
// @Summary  获取课程的学习进度
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    courseId path int true "课程ID"
// @Success  200  {object} response.Response{data=map[uint]model.LearningProgress} "获取成功"
// @Router   /progress/course/{courseId} [get]
func (a *ProgressApi) GetCourseProgress(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	progressMap, err := service.ServiceGroupApp.ProgressService.GetCourseProgress(userID, uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(progressMap, "获取成功", c)
}

// GetLessonProgress 获取单个课时的学习进度
// @Tags     Progress
// @Summary  获取单个课时的学习进度
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    lessonId path int true "课时ID"
// @Success  200  {object} response.Response{data=model.LearningProgress} "获取成功"
// @Router   /progress/lesson/{lessonId} [get]
func (a *ProgressApi) GetLessonProgress(c *gin.Context) {
	lessonIDStr := c.Param("lessonId")
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课时ID", c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	progress, err := service.ServiceGroupApp.ProgressService.GetLessonProgress(userID, uint(lessonID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(progress, "获取成功", c)
}
