package api

// ApiGroup API组
type ApiGroup struct {
	CourseApi        CourseApi
	ChapterLessonApi ChapterLessonApi
	EnrollmentApi    EnrollmentApi
	ProgressApi      ProgressApi
	HomeworkApi      HomeworkApi
	CertificateApi   CertificateApi
	CommentApi       CommentApi
}

var ApiGroupApp = new(ApiGroup)
