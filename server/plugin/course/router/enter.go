package router

// RouterGroup 路由组
type RouterGroup struct {
	CourseRouter        CourseRouter
	ChapterLessonRouter ChapterLessonRouter
	EnrollmentRouter    EnrollmentRouter
	ProgressRouter      ProgressRouter
	HomeworkRouter      HomeworkRouter
	CertificateRouter   CertificateRouter
	CommentRouter       CommentRouter
}

var RouterGroupApp = new(RouterGroup)
