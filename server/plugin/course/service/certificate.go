package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"gorm.io/gorm"
)

// CertificateService 证书服务
type CertificateService struct{}

// GenerateCertificate 生成证书
func (s *CertificateService) GenerateCertificate(userID, courseID uint) (model.Certificate, error) {
	// 检查用户是否报名了该课程
	var enrollment model.Enrollment
	err := global.GVA_DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Certificate{}, errors.New("未报名该课程")
		}
		return model.Certificate{}, err
	}

	// 检查是否已经生成证书
	var existingCertificate model.Certificate
	err = global.GVA_DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existingCertificate).Error
	if err == nil {
		return existingCertificate, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Certificate{}, err
	}

	// 检查课程是否存在
	var course model.Course
	err = global.GVA_DB.First(&course, courseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Certificate{}, errors.New("课程不存在")
		}
		return model.Certificate{}, err
	}

	// 检查用户的学习进度
	var lessons []model.Lesson
	err = global.GVA_DB.Joins("JOIN chapters ON chapters.id = lessons.chapter_id").Where("chapters.course_id = ?", courseID).Find(&lessons).Error
	if err != nil {
		return model.Certificate{}, err
	}

	if len(lessons) == 0 {
		return model.Certificate{}, errors.New("课程无内容")
	}

	// 检查所有课时是否已完成
	var completedCount int64
	err = global.GVA_DB.Model(&model.LearningProgress{}).Where("user_id = ? AND course_id = ? AND is_completed = ?", userID, courseID, true).Count(&completedCount).Error
	if err != nil {
		return model.Certificate{}, err
	}

	if int(completedCount) < len(lessons) {
		return model.Certificate{}, errors.New("未完成所有课程内容")
	}

	// 生成证书（这里简化处理，实际项目中可能需要生成PDF证书）
	certificateURL := "/certificates/" + course.Title + "_" + string(userID) + ".pdf"

	certificate := model.Certificate{
		UserID:         userID,
		CourseID:       courseID,
		CertificateURL: certificateURL,
	}

	err = global.GVA_DB.Create(&certificate).Error
	if err != nil {
		return model.Certificate{}, err
	}

	// 更新报名状态为已完成
	err = global.GVA_DB.Model(&enrollment).Update("status", "completed").Error
	if err != nil {
		return model.Certificate{}, err
	}

	return certificate, nil
}

// GetUserCertificates 获取用户的证书列表
func (s *CertificateService) GetUserCertificates(userID uint) ([]model.Certificate, error) {
	var certificates []model.Certificate

	err := global.GVA_DB.Where("user_id = ?", userID).Find(&certificates).Error
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

// GetCertificateByID 根据ID获取证书
func (s *CertificateService) GetCertificateByID(id uint) (model.Certificate, error) {
	var certificate model.Certificate
	err := global.GVA_DB.First(&certificate, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Certificate{}, errors.New("证书不存在")
		}
		return model.Certificate{}, err
	}

	return certificate, nil
}
