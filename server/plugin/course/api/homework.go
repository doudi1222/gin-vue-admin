package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// HomeworkApi 作业API
type HomeworkApi struct{}

// CreateHomework 创建作业
// @Tags     Homework
// @Summary  创建作业
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.CreateHomeworkRequest true "创建作业参数"
// @Success  200  {object} response.Response{data=model.Homework} "创建成功"
// @Router   /homework [post]
func (a *HomeworkApi) CreateHomework(c *gin.Context) {
	var req request.CreateHomeworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	homework, err := service.ServiceGroupApp.HomeworkService.CreateHomework(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(homework, "创建成功", c)
}

// UpdateHomework 更新作业
// @Tags     Homework
// @Summary  更新作业
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UpdateHomeworkRequest true "更新作业参数"
// @Success  200  {object} response.Response{data=model.Homework} "更新成功"
// @Router   /homework [put]
func (a *HomeworkApi) UpdateHomework(c *gin.Context) {
	var req request.UpdateHomeworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	homework, err := service.ServiceGroupApp.HomeworkService.UpdateHomework(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(homework, "更新成功", c)
}

// DeleteHomework 删除作业
// @Tags     Homework
// @Summary  删除作业
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "作业ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /homework/{id} [delete]
func (a *HomeworkApi) DeleteHomework(c *gin.Context) {
	id := c.Param("id")
	homeworkID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的作业ID", c)
		return
	}

	err = service.ServiceGroupApp.HomeworkService.DeleteHomework(uint(homeworkID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetCourseHomeworks 获取课程的作业列表
// @Tags     Homework
// @Summary  获取课程的作业列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    courseId path int true "课程ID"
// @Success  200  {object} response.Response{data=[]model.Homework} "获取成功"
// @Router   /homework/course/{courseId} [get]
func (a *HomeworkApi) GetCourseHomeworks(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	homeworks, err := service.ServiceGroupApp.HomeworkService.GetCourseHomeworks(uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(homeworks, "获取成功", c)
}

// SubmitHomework 提交作业
// @Tags     Homework
// @Summary  提交作业
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.SubmitHomeworkRequest true "提交作业参数"
// @Success  200  {object} response.Response{data=model.HomeworkSubmission} "提交成功"
// @Router   /homework/submit [post]
func (a *HomeworkApi) SubmitHomework(c *gin.Context) {
	var req request.SubmitHomeworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取
	req.UserID = userID

	submission, err := service.ServiceGroupApp.HomeworkService.SubmitHomework(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(submission, "提交成功", c)
}

// GradeHomework 批改作业
// @Tags     Homework
// @Summary  批改作业
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GradeHomeworkRequest true "批改作业参数"
// @Success  200  {object} response.Response{data=model.HomeworkSubmission} "批改成功"
// @Router   /homework/grade [put]
func (a *HomeworkApi) GradeHomework(c *gin.Context) {
	var req request.GradeHomeworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	submission, err := service.ServiceGroupApp.HomeworkService.GradeHomework(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(submission, "批改成功", c)
}

// GetHomeworkSubmissions 获取作业的提交列表
// @Tags     Homework
// @Summary  获取作业的提交列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    homeworkId path int true "作业ID"
// @Success  200  {object} response.Response{data=[]model.HomeworkSubmission} "获取成功"
// @Router   /homework/submissions/{homeworkId} [get]
func (a *HomeworkApi) GetHomeworkSubmissions(c *gin.Context) {
	homeworkIDStr := c.Param("homeworkId")
	homeworkID, err := strconv.ParseUint(homeworkIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的作业ID", c)
		return
	}

	submissions, err := service.ServiceGroupApp.HomeworkService.GetHomeworkSubmissions(uint(homeworkID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(submissions, "获取成功", c)
}

// GetUserHomeworkSubmission 获取用户的作业提交
// @Tags     Homework
// @Summary  获取用户的作业提交
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    homeworkId path int true "作业ID"
// @Success  200  {object} response.Response{data=model.HomeworkSubmission} "获取成功"
// @Router   /homework/my-submission/{homeworkId} [get]
func (a *HomeworkApi) GetUserHomeworkSubmission(c *gin.Context) {
	homeworkIDStr := c.Param("homeworkId")
	homeworkID, err := strconv.ParseUint(homeworkIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的作业ID", c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	submission, err := service.ServiceGroupApp.HomeworkService.GetUserHomeworkSubmission(uint(homeworkID), userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(submission, "获取成功", c)
}
