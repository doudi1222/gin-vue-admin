package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HomeworkSubmission 作业提交模型
type HomeworkSubmission struct {
	global.GVA_MODEL
	HomeworkID  uint      `json:"homeworkId" gorm:"column:homework_id;comment:作业ID;not null"`
	UserID      uint      `json:"userId" gorm:"column:user_id;comment:用户ID;not null"`
	Content     string    `json:"content" gorm:"column:content;comment:提交内容;type:text;not null"`
	FileURL     string    `json:"fileUrl" gorm:"column:file_url;comment:提交文件;size:255"`
	Grade       float64   `json:"grade" gorm:"column:grade;comment:评分;type:decimal(5,2)"`
	Feedback    string    `json:"feedback" gorm:"column:feedback;comment:反馈;type:text"`
	Status      string    `json:"status" gorm:"column:status;comment:状态;type:enum('submitted','graded');default:'submitted';not null"`
	SubmittedAt time.Time `json:"submittedAt" gorm:"column:submitted_at;comment:提交时间;not null;default:CURRENT_TIMESTAMP"`
	GradedAt    time.Time `json:"gradedAt" gorm:"column:graded_at;comment:批改时间"`
}

// TableName 设置表名
func (HomeworkSubmission) TableName() string {
	return "homework_submissions"
}
