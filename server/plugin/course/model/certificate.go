package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Certificate 证书模型
type Certificate struct {
	global.GVA_MODEL
	UserID          uint      `json:"userId" gorm:"column:user_id;comment:用户ID;not null"`
	CourseID        uint      `json:"courseId" gorm:"column:course_id;comment:课程ID;not null"`
	CertificateURL  string    `json:"certificateUrl" gorm:"column:certificate_url;comment:证书地址;size:255;not null"`
	IssuedAt        time.Time `json:"issuedAt" gorm:"column:issued_at;comment:颁发时间;not null;default:CURRENT_TIMESTAMP"`
}

// TableName 设置表名
func (Certificate) TableName() string {
	return "certificates"
}
