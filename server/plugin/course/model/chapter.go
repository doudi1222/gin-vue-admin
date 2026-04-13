package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Chapter 章节模型
type Chapter struct {
	global.GVA_MODEL
	CourseID uint     `json:"courseId" gorm:"column:course_id;comment:课程ID;not null"`
	Title    string   `json:"title" gorm:"column:title;comment:章节标题;size:100;not null"`
	Order    int      `json:"order" gorm:"column:order;comment:排序;default:0;not null"`
	Lessons  []Lesson `json:"lessons,omitempty" gorm:"foreignKey:ChapterID"`
}

// TableName 设置表名
func (Chapter) TableName() string {
	return "chapters"
}
