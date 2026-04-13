package request

// CourseSearch 课程搜索请求
type CourseSearch struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Category string `json:"category" form:"category"`
	Level    string `json:"level" form:"level"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   string `json:"status" form:"status"`
}

// CreateCourseRequest 创建课程请求
type CreateCourseRequest struct {
	Title         string  `json:"title" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Cover         string  `json:"cover" binding:"required"`
	Price         float64 `json:"price" binding:"required,gte=0"`
	OriginalPrice float64 `json:"originalPrice" binding:"required,gte=0"`
	Level         string  `json:"level" binding:"required,oneof=beginner intermediate advanced"`
	Category      string  `json:"category" binding:"required"`
	TeacherID     uint    `json:"teacherId" binding:"required"`
	TeacherName   string  `json:"teacherName" binding:"required"`
}

// UpdateCourseRequest 更新课程请求
type UpdateCourseRequest struct {
	ID            uint    `json:"id" binding:"required"`
	Title         string  `json:"title" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Cover         string  `json:"cover" binding:"required"`
	Price         float64 `json:"price" binding:"required,gte=0"`
	OriginalPrice float64 `json:"originalPrice" binding:"required,gte=0"`
	Level         string  `json:"level" binding:"required,oneof=beginner intermediate advanced"`
	Category      string  `json:"category" binding:"required"`
	Status        string  `json:"status" binding:"required,oneof=draft published archived"`
}

// EnrollCourseRequest 报名课程请求
type EnrollCourseRequest struct {
	CourseID uint `json:"courseId" binding:"required"`
}
