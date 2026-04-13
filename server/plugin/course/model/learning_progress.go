package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LearningProgress 学习进度模型
type LearningProgress struct {
	global.GVA_MODEL
	UserID           uint   `json:"userId" gorm:"column:user_id;comment:用户ID;not null"`
	CourseID         uint   `json:"courseId" gorm:"column:course_id;comment:课程ID;not null"`
	LessonID         uint   `json:"lessonId" gorm:"column:lesson_id;comment:课时ID;not null"`
	WatchedDuration  int    `json:"watchedDuration" gorm:"column:watched_duration;comment:已观看时长（秒）;default:0;not null"`
	IsCompleted      bool   `json:"isCompleted" gorm:"column:is_completed;comment:是否完成;default:false;not null"`
}

// TableName 设置表名
func (LearningProgress) TableName() string {
	return "learning_progress"
}
