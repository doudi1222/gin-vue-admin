package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/model/request"
	"gorm.io/gorm"
)

// CommentService 评论服务
type CommentService struct{}

// CreateComment 创建评论
func (s *CommentService) CreateComment(req request.CreateCommentRequest) (model.Comment, error) {
	// 检查课程是否存在
	var course model.Course
	err := global.GVA_DB.First(&course, req.CourseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Comment{}, errors.New("课程不存在")
		}
		return model.Comment{}, err
	}

	// 检查用户是否报名了该课程
	var enrollment model.Enrollment
	err = global.GVA_DB.Where("user_id = ? AND course_id = ?", req.UserID, req.CourseID).First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Comment{}, errors.New("未报名该课程，无法评论")
		}
		return model.Comment{}, err
	}

	// 如果是回复评论，检查父评论是否存在
	if req.ParentID != nil {
		var parentComment model.Comment
		err = global.GVA_DB.First(&parentComment, *req.ParentID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return model.Comment{}, errors.New("父评论不存在")
			}
			return model.Comment{}, err
		}

		// 检查父评论是否属于同一课程
		if parentComment.CourseID != req.CourseID {
			return model.Comment{}, errors.New("父评论不属于该课程")
		}
	}

	comment := model.Comment{
		CourseID: req.CourseID,
		UserID:   req.UserID,
		Content:  req.Content,
		Rating:   req.Rating,
		ParentID: req.ParentID,
	}

	err = global.GVA_DB.Create(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}

	// 更新课程的评分
	s.updateCourseRating(req.CourseID)

	return comment, nil
}

// GetCourseComments 获取课程的评论列表
func (s *CommentService) GetCourseComments(courseID uint) ([]model.Comment, error) {
	var comments []model.Comment

	err := global.GVA_DB.Where("course_id = ? AND parent_id IS NULL", courseID).Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	// 获取每个评论的回复
	for i := range comments {
		var replies []model.Comment
		err := global.GVA_DB.Where("parent_id = ?", comments[i].ID).Order("created_at ASC").Find(&replies).Error
		if err != nil {
			return nil, err
		}
		// 这里可以将回复添加到评论中，但需要在模型中添加相应字段
	}

	return comments, nil
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(id uint) error {
	// 检查评论是否存在
	var comment model.Comment
	err := global.GVA_DB.First(&comment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 删除评论的回复
	if err := tx.Where("parent_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除评论本身
	if err := tx.Delete(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新课程的评分
	if err := tx.Commit().Error; err != nil {
		return err
	}

	s.updateCourseRating(comment.CourseID)

	return nil
}

// updateCourseRating 更新课程的评分
func (s *CommentService) updateCourseRating(courseID uint) error {
	// 计算平均评分
	type Result struct {
		AvgRating float64
	}
	var result Result

	err := global.GVA_DB.Model(&model.Comment{}).Select("AVG(rating) as avg_rating").Where("course_id = ? AND parent_id IS NULL", courseID).Scan(&result).Error
	if err != nil {
		return err
	}

	// 更新课程评分
	err = global.GVA_DB.Model(&model.Course{}).Where("id = ?", courseID).Update("rating", result.AvgRating).Error
	if err != nil {
		return err
	}

	return nil
}
