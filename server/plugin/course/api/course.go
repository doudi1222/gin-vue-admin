package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// CourseApi 课程API
type CourseApi struct{}

// GetCourseList 获取课程列表
// @Tags     Course
// @Summary  获取课程列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data query request.CourseSearch true "课程搜索参数"
// @Success  200  {object} response.Response{data=response.PageResult{list=[]model.Course,total=int64}} "获取成功"
// @Router   /course/list [get]
func (a *CourseApi) GetCourseList(c *gin.Context) {
	var req request.CourseSearch
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	courses, total, err := service.ServiceGroupApp.CourseService.GetCourseList(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:  courses,
		Total: total,
	}, "获取成功", c)
}

// GetCourseByID 根据ID获取课程详情
// @Tags     Course
// @Summary  根据ID获取课程详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "课程ID"
// @Success  200  {object} response.Response{data=model.Course} "获取成功"
// @Router   /course/{id} [get]
func (a *CourseApi) GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	course, err := service.ServiceGroupApp.CourseService.GetCourseByID(uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(course, "获取成功", c)
}

// CreateCourse 创建课程
// @Tags     Course
// @Summary  创建课程
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.CreateCourseRequest true "创建课程参数"
// @Success  200  {object} response.Response{data=model.Course} "创建成功"
// @Router   /course [post]
func (a *CourseApi) CreateCourse(c *gin.Context) {
	var req request.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	course, err := service.ServiceGroupApp.CourseService.CreateCourse(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(course, "创建成功", c)
}

// UpdateCourse 更新课程
// @Tags     Course
// @Summary  更新课程
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.UpdateCourseRequest true "更新课程参数"
// @Success  200  {object} response.Response{data=model.Course} "更新成功"
// @Router   /course [put]
func (a *CourseApi) UpdateCourse(c *gin.Context) {
	var req request.UpdateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	course, err := service.ServiceGroupApp.CourseService.UpdateCourse(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(course, "更新成功", c)
}

// DeleteCourse 删除课程
// @Tags     Course
// @Summary  删除课程
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "课程ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /course/{id} [delete]
func (a *CourseApi) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	err = service.ServiceGroupApp.CourseService.DeleteCourse(uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
