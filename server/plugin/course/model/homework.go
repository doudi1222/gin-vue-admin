package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Homework 作业模型
type Homework struct {
	global.GVA_MODEL
	CourseID    uint      `json:"courseId" gorm:"column:course_id;comment:课程ID;not null"`
	Title       string    `json:"title" gorm:"column:title;comment:作业标题;size:100;not null"`
	Description string    `json:"description" gorm:"column:description;comment:作业描述;type:text;not null"`
	Deadline    time.Time `json:"deadline" gorm:"column:deadline;comment:截止时间"`
}

// TableName 设置表名
func (Homework) TableName() string {
	return "homeworks"
}
