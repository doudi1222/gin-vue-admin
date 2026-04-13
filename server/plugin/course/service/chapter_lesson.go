package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"gorm.io/gorm"
)

// ChapterLessonService 章节和课时服务
type ChapterLessonService struct{}

// CreateChapter 创建章节
func (s *ChapterLessonService) CreateChapter(req request.CreateChapterRequest) (model.Chapter, error) {
	// 检查课程是否存在
	var course model.Course
	err := global.GVA_DB.First(&course, req.CourseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Chapter{}, errors.New("课程不存在")
		}
		return model.Chapter{}, err
	}

	chapter := model.Chapter{
		CourseID: req.CourseID,
		Title:    req.Title,
		Order:    req.Order,
	}

	err = global.GVA_DB.Create(&chapter).Error
	if err != nil {
		return model.Chapter{}, err
	}

	return chapter, nil
}

// UpdateChapter 更新章节
func (s *ChapterLessonService) UpdateChapter(req request.UpdateChapterRequest) (model.Chapter, error) {
	var chapter model.Chapter
	err := global.GVA_DB.First(&chapter, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Chapter{}, errors.New("章节不存在")
		}
		return model.Chapter{}, err
	}

	chapter.Title = req.Title
	chapter.Order = req.Order

	err = global.GVA_DB.Save(&chapter).Error
	if err != nil {
		return model.Chapter{}, err
	}

	return chapter, nil
}

// DeleteChapter 删除章节
func (s *ChapterLessonService) DeleteChapter(id uint) error {
	// 检查章节是否存在
	var chapter model.Chapter
	err := global.GVA_DB.First(&chapter, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("章节不存在")
		}
		return err
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 删除章节的课时
	if err := tx.Where("chapter_id = ?", id).Delete(&model.Lesson{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除章节本身
	if err := tx.Delete(&chapter).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// CreateLesson 创建课时
func (s *ChapterLessonService) CreateLesson(req request.CreateLessonRequest) (model.Lesson, error) {
	// 检查章节是否存在
	var chapter model.Chapter
	err := global.GVA_DB.First(&chapter, req.ChapterID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Lesson{}, errors.New("章节不存在")
		}
		return model.Lesson{}, err
	}

	lesson := model.Lesson{
		ChapterID: req.ChapterID,
		Title:     req.Title,
		VideoURL:  req.VideoURL,
		Duration:  req.Duration,
		Content:   req.Content,
		Order:     req.Order,
	}

	err = global.GVA_DB.Create(&lesson).Error
	if err != nil {
		return model.Lesson{}, err
	}

	return lesson, nil
}

// UpdateLesson 更新课时
func (s *ChapterLessonService) UpdateLesson(req request.UpdateLessonRequest) (model.Lesson, error) {
	var lesson model.Lesson
	err := global.GVA_DB.First(&lesson, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Lesson{}, errors.New("课时不存在")
		}
		return model.Lesson{}, err
	}

	lesson.Title = req.Title
	lesson.VideoURL = req.VideoURL
	lesson.Duration = req.Duration
	lesson.Content = req.Content
	lesson.Order = req.Order

	err = global.GVA_DB.Save(&lesson).Error
	if err != nil {
		return model.Lesson{}, err
	}

	return lesson, nil
}

// DeleteLesson 删除课时
func (s *ChapterLessonService) DeleteLesson(id uint) error {
	// 检查课时是否存在
	var lesson model.Lesson
	err := global.GVA_DB.First(&lesson, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("课时不存在")
		}
		return err
	}

	// 删除课时
	err = global.GVA_DB.Delete(&lesson).Error
	if err != nil {
		return err
	}

	// 删除相关的学习进度
	err = global.GVA_DB.Where("lesson_id = ?", id).Delete(&model.LearningProgress{}).Error
	if err != nil {
		return err
	}

	return nil
}
