package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/service"
	"github.com/gin-gonic/gin"
)

// CertificateApi 证书API
type CertificateApi struct{}

// GenerateCertificate 生成证书
// @Tags     Certificate
// @Summary  生成证书
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    courseId path int true "课程ID"
// @Success  200  {object} response.Response{data=model.Certificate} "生成成功"
// @Router   /certificate/generate/{courseId} [post]
func (a *CertificateApi) GenerateCertificate(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的课程ID", c)
		return
	}

	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	certificate, err := service.ServiceGroupApp.CertificateService.GenerateCertificate(userID, uint(courseID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(certificate, "生成成功", c)
}

// GetUserCertificates 获取用户的证书列表
// @Tags     Certificate
// @Summary  获取用户的证书列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Certificate} "获取成功"
// @Router   /certificate/my-certificates [get]
func (a *CertificateApi) GetUserCertificates(c *gin.Context) {
	// 从上下文获取用户ID（实际项目中应该从JWT token中获取）
	userID := uint(1) // 这里简化处理，实际应该从上下文获取

	certificates, err := service.ServiceGroupApp.CertificateService.GetUserCertificates(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(certificates, "获取成功", c)
}

// GetCertificateByID 根据ID获取证书
// @Tags     Certificate
// @Summary  根据ID获取证书
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path int true "证书ID"
// @Success  200  {object} response.Response{data=model.Certificate} "获取成功"
// @Router   /certificate/{id} [get]
func (a *CertificateApi) GetCertificateByID(c *gin.Context) {
	id := c.Param("id")
	certificateID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的证书ID", c)
		return
	}

	certificate, err := service.ServiceGroupApp.CertificateService.GetCertificateByID(uint(certificateID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(certificate, "获取成功", c)
}
