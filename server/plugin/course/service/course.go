package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"gorm.io/gorm"
)

// CourseService 课程服务
type CourseService struct{}

// GetCourseList 获取课程列表
func (s *CourseService) GetCourseList(req request.CourseSearch) ([]model.Course, int64, error) {
	var courses []model.Course
	var total int64

	db := global.GVA_DB.Model(&model.Course{})

	// 筛选条件
	if req.Category != "" {
		db = db.Where("category = ?", req.Category)
	}
	if req.Level != "" {
		db = db.Where("level = ?", req.Level)
	}
	if req.Keyword != "" {
		db = db.Where("title LIKE ? OR description LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	// 计算总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (req.Page - 1) * req.PageSize
	err = db.Offset(offset).Limit(req.PageSize).Order("created_at DESC").Find(&courses).Error
	if err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

// GetCourseByID 根据ID获取课程详情
func (s *CourseService) GetCourseByID(id uint) (model.Course, error) {
	var course model.Course
	err := global.GVA_DB.Preload("Chapters", func(db *gorm.DB) *gorm.DB {
		return db.Order("order ASC").Preload("Lessons", func(db *gorm.DB) *gorm.DB {
			return db.Order("order ASC")
		})
	}).First(&course, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Course{}, errors.New("课程不存在")
		}
		return model.Course{}, err
	}

	return course, nil
}

// CreateCourse 创建课程
func (s *CourseService) CreateCourse(req request.CreateCourseRequest) (model.Course, error) {
	course := model.Course{
		Title:         req.Title,
		Description:   req.Description,
		Cover:         req.Cover,
		Price:         req.Price,
		OriginalPrice: req.OriginalPrice,
		Level:         req.Level,
		Category:      req.Category,
		TeacherID:     req.TeacherID,
		TeacherName:   req.TeacherName,
		Status:        "draft",
	}

	err := global.GVA_DB.Create(&course).Error
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

// UpdateCourse 更新课程
func (s *CourseService) UpdateCourse(req request.UpdateCourseRequest) (model.Course, error) {
	var course model.Course
	err := global.GVA_DB.First(&course, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Course{}, errors.New("课程不存在")
		}
		return model.Course{}, err
	}

	course.Title = req.Title
	course.Description = req.Description
	course.Cover = req.Cover
	course.Price = req.Price
	course.OriginalPrice = req.OriginalPrice
	course.Level = req.Level
	course.Category = req.Category
	course.Status = req.Status

	err = global.GVA_DB.Save(&course).Error
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

// DeleteCourse 删除课程
func (s *CourseService) DeleteCourse(id uint) error {
	// 检查课程是否存在
	var course model.Course
	err := global.GVA_DB.First(&course, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("课程不存在")
		}
		return err
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 删除课程的章节和课时
	var chapters []model.Chapter
	if err := tx.Where("course_id = ?", id).Find(&chapters).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, chapter := range chapters {
		if err := tx.Where("chapter_id = ?", chapter.ID).Delete(&model.Lesson{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Where("course_id = ?", id).Delete(&model.Chapter{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除课程的报名记录
	if err := tx.Where("course_id = ?", id).Delete(&model.Enrollment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除课程的学习进度
	if err := tx.Where("course_id = ?", id).Delete(&model.LearningProgress{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除课程的作业
	var homeworks []model.Homework
	if err := tx.Where("course_id = ?", id).Find(&homeworks).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, homework := range homeworks {
		if err := tx.Where("homework_id = ?", homework.ID).Delete(&model.HomeworkSubmission{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Where("course_id = ?", id).Delete(&model.Homework{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除课程的证书
	if err := tx.Where("course_id = ?", id).Delete(&model.Certificate{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除课程的评论
	if err := tx.Where("course_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除课程本身
	if err := tx.Delete(&course).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
