package request

// CreateChapterRequest 创建章节请求
type CreateChapterRequest struct {
	CourseID uint   `json:"courseId" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Order    int    `json:"order" binding:"required,gte=0"`
}

// UpdateChapterRequest 更新章节请求
type UpdateChapterRequest struct {
	ID     uint   `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Order  int    `json:"order" binding:"required,gte=0"`
}

// CreateLessonRequest 创建课时请求
type CreateLessonRequest struct {
	ChapterID uint   `json:"chapterId" binding:"required"`
	Title     string `json:"title" binding:"required"`
	VideoURL  string `json:"videoUrl" binding:"required"`
	Duration  int    `json:"duration" binding:"required,gte=0"`
	Content   string `json:"content"`
	Order     int    `json:"order" binding:"required,gte=0"`
}

// UpdateLessonRequest 更新课时请求
type UpdateLessonRequest struct {
	ID        uint   `json:"id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	VideoURL  string `json:"videoUrl" binding:"required"`
	Duration  int    `json:"duration" binding:"required,gte=0"`
	Content   string `json:"content"`
	Order     int    `json:"order" binding:"required,gte=0"`
}
