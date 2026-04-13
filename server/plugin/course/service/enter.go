package service

// ServiceGroup 服务组
type ServiceGroup struct {
	CourseService         CourseService
	ChapterLessonService  ChapterLessonService
	EnrollmentService     EnrollmentService
	ProgressService       ProgressService
	HomeworkService       HomeworkService
	CertificateService    CertificateService
	CommentService        CommentService
}

var ServiceGroupApp = new(ServiceGroup)
