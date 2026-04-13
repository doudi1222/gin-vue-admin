package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"gorm.io/gorm"
)

// ProgressService 学习进度服务
type ProgressService struct{}

// UpdateProgress 更新学习进度
func (s *ProgressService) UpdateProgress(req request.UpdateProgressRequest) (model.LearningProgress, error) {
	// 检查用户是否报名了该课程
	var enrollment model.Enrollment
	err := global.GVA_DB.Where("user_id = ? AND course_id = ?", req.UserID, req.CourseID).First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.LearningProgress{}, errors.New("未报名该课程")
		}
		return model.LearningProgress{}, err
	}

	// 检查课时是否存在
	var lesson model.Lesson
	err = global.GVA_DB.First(&lesson, req.LessonID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.LearningProgress{}, errors.New("课时不存在")
		}
		return model.LearningProgress{}, err
	}

	// 查找或创建学习进度记录
	var progress model.LearningProgress
	err = global.GVA_DB.Where("user_id = ? AND lesson_id = ?", req.UserID, req.LessonID).First(&progress).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新的学习进度记录
			progress = model.LearningProgress{
				UserID:          req.UserID,
				CourseID:        req.CourseID,
				LessonID:        req.LessonID,
				WatchedDuration: req.WatchedDuration,
				IsCompleted:     req.IsCompleted,
			}
			err = global.GVA_DB.Create(&progress).Error
		} else {
			return model.LearningProgress{}, err
		}
	} else {
		// 更新现有记录
		progress.WatchedDuration = req.WatchedDuration
		progress.IsCompleted = req.IsCompleted
		err = global.GVA_DB.Save(&progress).Error
	}

	if err != nil {
		return model.LearningProgress{}, err
	}

	return progress, nil
}

// GetCourseProgress 获取课程的学习进度
func (s *ProgressService) GetCourseProgress(userID, courseID uint) (map[uint]model.LearningProgress, error) {
	var progresses []model.LearningProgress

	err := global.GVA_DB.Where("user_id = ? AND course_id = ?", userID, courseID).Find(&progresses).Error
	if err != nil {
		return nil, err
	}

	// 将进度记录转换为map，方便前端使用
	progressMap := make(map[uint]model.LearningProgress)
	for _, progress := range progresses {
		progressMap[progress.LessonID] = progress
	}

	return progressMap, nil
}

// GetLessonProgress 获取单个课时的学习进度
func (s *ProgressService) GetLessonProgress(userID, lessonID uint) (model.LearningProgress, error) {
	var progress model.LearningProgress
	err := global.GVA_DB.Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&progress).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 返回空进度
			return model.LearningProgress{
				UserID:          userID,
				LessonID:        lessonID,
				WatchedDuration: 0,
				IsCompleted:     false,
			}, nil
		}
		return model.LearningProgress{}, err
	}

	return progress, nil
}
