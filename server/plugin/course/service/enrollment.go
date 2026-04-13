package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"gorm.io/gorm"
)

// EnrollmentService 报名服务
type EnrollmentService struct{}

// EnrollCourse 报名课程
func (s *EnrollmentService) EnrollCourse(userID, courseID uint) (model.Enrollment, error) {
	// 检查课程是否存在
	var course model.Course
	err := global.GVA_DB.First(&course, courseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Enrollment{}, errors.New("课程不存在")
		}
		return model.Enrollment{}, err
	}

	// 检查课程状态
	if course.Status != "published" {
		return model.Enrollment{}, errors.New("课程未发布")
	}

	// 检查是否已经报名
	var existingEnrollment model.Enrollment
	err = global.GVA_DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existingEnrollment).Error
	if err == nil {
		return model.Enrollment{}, errors.New("已经报名该课程")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Enrollment{}, err
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 创建报名记录
	enrollment := model.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "active",
	}

	if err := tx.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return model.Enrollment{}, err
	}

	// 更新课程的学生人数
	if err := tx.Model(&course).Update("student_count", course.StudentCount+1).Error; err != nil {
		tx.Rollback()
		return model.Enrollment{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return model.Enrollment{}, err
	}

	return enrollment, nil
}

// GetUserCourses 获取用户的课程列表
func (s *EnrollmentService) GetUserCourses(userID uint) ([]model.Course, error) {
	var courses []model.Course

	err := global.GVA_DB.Joins("JOIN enrollments ON enrollments.course_id = courses.id").Where("enrollments.user_id = ?", userID).Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

// GetCourseStudents 获取课程的学生列表
func (s *EnrollmentService) GetCourseStudents(courseID uint) ([]uint, error) {
	var enrollments []model.Enrollment

	err := global.GVA_DB.Where("course_id = ?", courseID).Find(&enrollments).Error
	if err != nil {
		return nil, err
	}

	studentIDs := make([]uint, len(enrollments))
	for i, enrollment := range enrollments {
		studentIDs[i] = enrollment.UserID
	}

	return studentIDs, nil
}

// CancelEnrollment 取消报名
func (s *EnrollmentService) CancelEnrollment(userID, courseID uint) error {
	// 检查报名记录是否存在
	var enrollment model.Enrollment
	err := global.GVA_DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("未报名该课程")
		}
		return err
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 删除报名记录
	if err := tx.Delete(&enrollment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新课程的学生人数
	var course model.Course
	if err := tx.First(&course, courseID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if course.StudentCount > 0 {
		if err := tx.Model(&course).Update("student_count", course.StudentCount-1).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
