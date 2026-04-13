package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Course 课程模型
type Course struct {
	global.GVA_MODEL
	Title          string  `json:"title" gorm:"column:title;comment:课程标题;size:100;not null"`
	Description    string  `json:"description" gorm:"column:description;comment:课程描述;type:text;not null"`
	Cover          string  `json:"cover" gorm:"column:cover;comment:封面图片;size:255;not null"`
	Price          float64 `json:"price" gorm:"column:price;comment:价格;type:decimal(10,2);default:0.00;not null"`
	OriginalPrice  float64 `json:"originalPrice" gorm:"column:original_price;comment:原价;type:decimal(10,2);default:0.00;not null"`
	Level          string  `json:"level" gorm:"column:level;comment:难度级别;type:enum('beginner','intermediate','advanced');default:'beginner';not null"`
	Category       string  `json:"category" gorm:"column:category;comment:分类;size:50;not null"`
	TeacherID      uint    `json:"teacherId" gorm:"column:teacher_id;comment:教师ID;not null"`
	TeacherName    string  `json:"teacherName" gorm:"column:teacher_name;comment:教师姓名;size:50;not null"`
	Status         string  `json:"status" gorm:"column:status;comment:状态;type:enum('draft','published','archived');default:'draft';not null"`
	StudentCount   int     `json:"studentCount" gorm:"column:student_count;comment:学生人数;default:0;not null"`
	Rating         float64 `json:"rating" gorm:"column:rating;comment:评分;type:decimal(3,2);default:0.00;not null"`
	Chapters       []Chapter `json:"chapters,omitempty" gorm:"foreignKey:CourseID"`
}

// TableName 设置表名
func (Course) TableName() string {
	return "courses"
}
