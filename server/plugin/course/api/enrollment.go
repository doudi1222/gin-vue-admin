package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// EnrollmentApi 报名API
type EnrollmentApi struct{}

// EnrollCourse 报名课程
// @Tags     Enrollment
// @Summary  报名课程
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.EnrollCourseRequest true "报名课程参数"
// @Success  200  {object} response.Response{data=model.Enrollment} "报名成功"
// @Router   /enrollment [post]
func (a *EnrollmentApi) EnrollCourse(c *gin.Context) {
	var req request.EnrollCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	enrollment, err := service.ServiceGroupApp.EnrollmentService.EnrollCourse(userID, req.CourseID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(enrollment, "报名成功", c)
}

// GetUserCourses 获取用户的课程列表
// @Tags     Enrollment
// @Summary  获取用户的课程列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Course} "获取成功"
// @Router   /enrollment/my-courses [get]
func (a *EnrollmentApi) GetUserCourses(c *gin.Context) {
	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	courses, err := service.ServiceGroupApp.EnrollmentService.GetUserCourses(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(courses, "获取成功", c)
}

// GetCourseStudents 获取课程的学生列表
// @Tags     Enrollment
// @Summary  获取课程的学生列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    courseId path int true "课程ID"
// @Success  200  {object} response.Response{data=[]uint} "获取成功"
// @Router   /enrollment/course-students/{courseId} [get]
func (a *EnrollmentApi) GetCourseStudents(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	studentIDs, err := service.ServiceGroupApp.EnrollmentService.GetCourseStudents(uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(studentIDs, "获取成功", c)
}

// CancelEnrollment 取消报名
// @Tags     Enrollment
// @Summary  取消报名
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.EnrollCourseRequest true "取消报名参数"
// @Success  200  {object} response.Response{msg=string} "取消成功"
// @Router   /enrollment/cancel [post]
func (a *EnrollmentApi) CancelEnrollment(c *gin.Context) {
	var req request.EnrollCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	err := service.ServiceGroupApp.EnrollmentService.CancelEnrollment(userID, req.CourseID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("取消成功", c)
}
