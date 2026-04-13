package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Lesson 课时模型
type Lesson struct {
	global.GVA_MODEL
	ChapterID     uint   `json:"chapterId" gorm:"column:chapter_id;comment:章节ID;not null"`
	Title         string `json:"title" gorm:"column:title;comment:课时标题;size:100;not null"`
	VideoURL      string `json:"videoUrl" gorm:"column:video_url;comment:视频地址;size:255;not null"`
	Duration      int    `json:"duration" gorm:"column:duration;comment:时长（秒）;default:0;not null"`
	Content       string `json:"content" gorm:"column:content;comment:内容;type:text"`
	Order         int    `json:"order" gorm:"column:order;comment:排序;default:0;not null"`
}

// TableName 设置表名
func (Lesson) TableName() string {
	return "lessons"
}
