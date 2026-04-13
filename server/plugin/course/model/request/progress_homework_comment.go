package request

// UpdateProgressRequest 更新学习进度请求
type UpdateProgressRequest struct {
	UserID          uint `json:"userId" binding:"required"`
	CourseID        uint `json:"courseId" binding:"required"`
	LessonID        uint `json:"lessonId" binding:"required"`
	WatchedDuration int  `json:"watchedDuration" binding:"required,gte=0"`
	IsCompleted     bool `json:"isCompleted"`
}

// CreateHomeworkRequest 创建作业请求
type CreateHomeworkRequest struct {
	CourseID    uint   `json:"courseId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline    string `json:"deadline"`
}

// UpdateHomeworkRequest 更新作业请求
type UpdateHomeworkRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline    string `json:"deadline"`
}

// SubmitHomeworkRequest 提交作业请求
type SubmitHomeworkRequest struct {
	HomeworkID uint   `json:"homeworkId" binding:"required"`
	UserID     uint   `json:"userId" binding:"required"`
	Content    string `json:"content" binding:"required"`
	FileURL    string `json:"fileUrl"`
}

// GradeHomeworkRequest 批改作业请求
type GradeHomeworkRequest struct {
	ID       uint    `json:"id" binding:"required"`
	Grade    float64 `json:"grade" binding:"required,gte=0,lte=100"`
	Feedback string  `json:"feedback"`
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	CourseID uint   `json:"courseId" binding:"required"`
	UserID   uint   `json:"userId" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Rating   int    `json:"rating" binding:"required,gte=1,lte=5"`
	ParentID *uint  `json:"parentId"`
}
