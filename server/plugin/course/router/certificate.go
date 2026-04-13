package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/api"
	"github.com/gin-gonic/gin"
)

// CertificateRouter 证书路由
type CertificateRouter struct{}

// InitCertificateRouter 初始化证书路由
func (r *CertificateRouter) InitCertificateRouter(router *gin.RouterGroup) {
	certificateRouter := router.Group("/certificate")
	{
		certificateRouter.POST("/generate/:courseId", api.ApiGroupApp.CertificateApi.GenerateCertificate)
		certificateRouter.GET("/my-certificates", api.ApiGroupApp.CertificateApi.GetUserCertificates)
		certificateRouter.GET("/:id", api.ApiGroupApp.CertificateApi.GetCertificateByID)
	}
}
