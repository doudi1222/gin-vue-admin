package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Enrollment 报名记录模型
type Enrollment struct {
	global.GVA_MODEL
	UserID    uint   `json:"userId" gorm:"column:user_id;comment:用户ID;not null"`
	CourseID  uint   `json:"courseId" gorm:"column:course_id;comment:课程ID;not null"`
	Status    string `json:"status" gorm:"column:status;comment:状态;type:enum('active','completed');default:'active';not null"`
}

// TableName 设置表名
func (Enrollment) TableName() string {
	return "enrollments"
}
