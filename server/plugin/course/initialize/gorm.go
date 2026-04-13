package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"gorm.io/gorm"
)

// InitDB 初始化数据库
func InitDB(db *gorm.DB) {
	// 自动迁移表结构
	err := db.AutoMigrate(
		&model.Course{},
		&model.Chapter{},
		&model.Lesson{},
		&model.Enrollment{},
		&model.LearningProgress{},
		&model.Homework{},
		&model.HomeworkSubmission{},
		&model.Certificate{},
		&model.Comment{},
	)

	if err != nil {
		global.GVA_LOG.Error("自动迁移失败", err)
		panic("自动迁移失败")
	}

	global.GVA_LOG.Info("数据库初始化成功")
}
