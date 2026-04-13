package service

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"gorm.io/gorm"
)

// HomeworkService 作业服务
type HomeworkService struct{}

// CreateHomework 创建作业
func (s *HomeworkService) CreateHomework(req request.CreateHomeworkRequest) (model.Homework, error) {
	// 检查课程是否存在
	var course model.Course
	err := global.GVA_DB.First(&course, req.CourseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Homework{}, errors.New("课程不存在")
		}
		return model.Homework{}, err
	}

	var deadline time.Time
	if req.Deadline != "" {
		deadline, err = time.Parse("2006-01-02 15:04:05", req.Deadline)
		if err != nil {
			return model.Homework{}, errors.New("截止时间格式错误")
		}
	}

	homework := model.Homework{
		CourseID:    req.CourseID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    deadline,
	}

	err = global.GVA_DB.Create(&homework).Error
	if err != nil {
		return model.Homework{}, err
	}

	return homework, nil
}

// UpdateHomework 更新作业
func (s *HomeworkService) UpdateHomework(req request.UpdateHomeworkRequest) (model.Homework, error) {
	var homework model.Homework
	err := global.GVA_DB.First(&homework, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Homework{}, errors.New("作业不存在")
		}
		return model.Homework{}, err
	}

	var deadline time.Time
	if req.Deadline != "" {
		deadline, err = time.Parse("2006-01-02 15:04:05", req.Deadline)
		if err != nil {
			return model.Homework{}, errors.New("截止时间格式错误")
		}
		homework.Deadline = deadline
	}

	homework.Title = req.Title
	homework.Description = req.Description

	err = global.GVA_DB.Save(&homework).Error
	if err != nil {
		return model.Homework{}, err
	}

	return homework, nil
}

// DeleteHomework 删除作业
func (s *HomeworkService) DeleteHomework(id uint) error {
	// 检查作业是否存在
	var homework model.Homework
	err := global.GVA_DB.First(&homework, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("作业不存在")
		}
		return err
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 删除作业的提交记录
	if err := tx.Where("homework_id = ?", id).Delete(&model.HomeworkSubmission{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除作业本身
	if err := tx.Delete(&homework).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetCourseHomeworks 获取课程的作业列表
func (s *HomeworkService) GetCourseHomeworks(courseID uint) ([]model.Homework, error) {
	var homeworks []model.Homework

	err := global.GVA_DB.Where("course_id = ?", courseID).Order("created_at DESC").Find(&homeworks).Error
	if err != nil {
		return nil, err
	}

	return homeworks, nil
}

// SubmitHomework 提交作业
func (s *HomeworkService) SubmitHomework(req request.SubmitHomeworkRequest) (model.HomeworkSubmission, error) {
	// 检查作业是否存在
	var homework model.Homework
	err := global.GVA_DB.First(&homework, req.HomeworkID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.HomeworkSubmission{}, errors.New("作业不存在")
		}
		return model.HomeworkSubmission{}, err
	}

	// 检查用户是否报名了该课程
	var enrollment model.Enrollment
	err = global.GVA_DB.Where("user_id = ? AND course_id = ?", req.UserID, homework.CourseID).First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.HomeworkSubmission{}, errors.New("未报名该课程")
		}
		return model.HomeworkSubmission{}, err
	}

	// 检查是否已经提交
	var existingSubmission model.HomeworkSubmission
	err = global.GVA_DB.Where("homework_id = ? AND user_id = ?", req.HomeworkID, req.UserID).First(&existingSubmission).Error
	if err == nil {
		// 更新提交
		existingSubmission.Content = req.Content
		existingSubmission.FileURL = req.FileURL
		existingSubmission.Status = "submitted"
		existingSubmission.SubmittedAt = time.Now()
		existingSubmission.Grade = 0
		existingSubmission.Feedback = ""
		existingSubmission.GradedAt = time.Time{}

		err = global.GVA_DB.Save(&existingSubmission).Error
		if err != nil {
			return model.HomeworkSubmission{}, err
		}

		return existingSubmission, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.HomeworkSubmission{}, err
	}

	// 创建新的提交记录
	submission := model.HomeworkSubmission{
		HomeworkID:  req.HomeworkID,
		UserID:      req.UserID,
		Content:     req.Content,
		FileURL:     req.FileURL,
		Status:      "submitted",
		SubmittedAt: time.Now(),
	}

	err = global.GVA_DB.Create(&submission).Error
	if err != nil {
		return model.HomeworkSubmission{}, err
	}

	return submission, nil
}

// GradeHomework 批改作业
func (s *HomeworkService) GradeHomework(req request.GradeHomeworkRequest) (model.HomeworkSubmission, error) {
	var submission model.HomeworkSubmission
	err := global.GVA_DB.First(&submission, req.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.HomeworkSubmission{}, errors.New("作业提交记录不存在")
		}
		return model.HomeworkSubmission{}, err
	}

	submission.Grade = req.Grade
	submission.Feedback = req.Feedback
	submission.Status = "graded"
	submission.GradedAt = time.Now()

	err = global.GVA_DB.Save(&submission).Error
	if err != nil {
		return model.HomeworkSubmission{}, err
	}

	return submission, nil
}

// GetHomeworkSubmissions 获取作业的提交列表
func (s *HomeworkService) GetHomeworkSubmissions(homeworkID uint) ([]model.HomeworkSubmission, error) {
	var submissions []model.HomeworkSubmission

	err := global.GVA_DB.Where("homework_id = ?", homeworkID).Order("submitted_at DESC").Find(&submissions).Error
	if err != nil {
		return nil, err
	}

	return submissions, nil
}

// GetUserHomeworkSubmission 获取用户的作业提交
func (s *HomeworkService) GetUserHomeworkSubmission(homeworkID, userID uint) (model.HomeworkSubmission, error) {
	var submission model.HomeworkSubmission
	err := global.GVA_DB.Where("homework_id = ? AND user_id = ?", homeworkID, userID).First(&submission).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.HomeworkSubmission{}, errors.New("未提交作业")
		}
		return model.HomeworkSubmission{}, err
	}

	return submission, nil
}
