package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// CommentApi 评论API
type CommentApi struct{}

// CreateComment 创建评论
// @Tags     Comment
// @Summary  创建评论
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.CreateCommentRequest true "创建评论参数"
// @Success  200  {object} response.Response{data=model.Comment} "创建成功"
// @Router   /comment [post]
func (a *CommentApi) CreateComment(c *gin.Context) {
	var req request.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取
	req.UserID = userID

	comment, err := service.ServiceGroupApp.CommentService.CreateComment(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(comment, "创建成功", c)
}

// GetCourseComments 获取课程的评论列表
// @Tags     Comment
// @Summary  获取课程的评论列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    courseId path int true "课程ID"
// @Success  200  {object} response.Response{data=[]model.Comment} "获取成功"
// @Router   /comment/course/{courseId} [get]
func (a *CommentApi) GetCourseComments(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	comments, err := service.ServiceGroupApp.CommentService.GetCourseComments(uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(comments, "获取成功", c)
}

// DeleteComment 删除评论
// @Tags     Comment
// @Summary  删除评论
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "评论ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /comment/{id} [delete]
func (a *CommentApi) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的评论ID", c)
		return
	}

	err = service.ServiceGroupApp.CommentService.DeleteComment(uint(commentID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
