package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Comment 评论模型
type Comment struct {
	global.GVA_MODEL
	CourseID  uint   `json:"courseId" gorm:"column:course_id;comment:课程ID;not null"`
	UserID    uint   `json:"userId" gorm:"column:user_id;comment:用户ID;not null"`
	Content   string `json:"content" gorm:"column:content;comment:评论内容;type:text;not null"`
	Rating    int    `json:"rating" gorm:"column:rating;comment:评分（1-5）;default:5;not null"`
	ParentID  *uint  `json:"parentId" gorm:"column:parent_id;comment:父评论ID"`
}

// TableName 设置表名
func (Comment) TableName() string {
	return "comments"
}
