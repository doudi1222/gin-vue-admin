package initialize

import (
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemSource "github.com/flipped-aurora/gin-vue-admin/server/source/system"
)

// InitMenu 初始化菜单
func InitMenu() {
	// 课程管理菜单
	courseMenu := systemModel.SysBaseMenu{
		MenuLevel:  1,
		ParentID:   0,
		Path:       "/course",
		Name:       "course",
		Hidden:     false,
		Component:  "layout",
		Title:      "课程管理",
		Icon:       "book",
		Sort:       10,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "book", Title: "课程管理"},
	}

	// 课程列表子菜单
	courseListMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "list",
		Name:       "courseList",
		Hidden:     false,
		Component:  "course/list",
		Title:      "课程列表",
		Icon:       "list",
		Sort:       1,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "list", Title: "课程列表"},
	}

	// 学习中心菜单
	learningMenu := systemModel.SysBaseMenu{
		MenuLevel:  1,
		ParentID:   0,
		Path:       "/learning",
		Name:       "learning",
		Hidden:     false,
		Component:  "layout",
		Title:      "学习中心",
		Icon:       "education",
		Sort:       11,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "education", Title: "学习中心"},
	}

	// 我的课程子菜单
	myCoursesMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "my-courses",
		Name:       "myCourses",
		Hidden:     false,
		Component:  "learning/my-courses",
		Title:      "我的课程",
		Icon:       "collection",
		Sort:       1,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "collection", Title: "我的课程"},
	}

	// 作业管理子菜单
	homeworkMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "homework",
		Name:       "homework",
		Hidden:     false,
		Component:  "learning/homework",
		Title:      "作业管理",
		Icon:       "document",
		Sort:       2,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "document", Title: "作业管理"},
	}

	// 证书管理子菜单
	certificateMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "certificate",
		Name:       "certificate",
		Hidden:     false,
		Component:  "learning/certificate",
		Title:      "证书管理",
		Icon:       "medal",
		Sort:       3,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "medal", Title: "证书管理"},
	}

	// 教师中心菜单
	teacherMenu := systemModel.SysBaseMenu{
		MenuLevel:  1,
		ParentID:   0,
		Path:       "/teacher",
		Name:       "teacher",
		Hidden:     false,
		Component:  "layout",
		Title:      "教师中心",
		Icon:       "user",
		Sort:       12,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "user", Title: "教师中心"},
	}

	// 课程管理子菜单（教师）
	teacherCourseMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "course",
		Name:       "teacherCourse",
		Hidden:     false,
		Component:  "teacher/course",
		Title:      "课程管理",
		Icon:       "edit",
		Sort:       1,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "edit", Title: "课程管理"},
	}

	// 学生管理子菜单
	studentMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "student",
		Name:       "student",
		Hidden:     false,
		Component:  "teacher/student",
		Title:      "学生管理",
		Icon:       "users",
		Sort:       2,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "users", Title: "学生管理"},
	}

	// 数据分析子菜单
	analyticsMenu := systemModel.SysBaseMenu{
		MenuLevel:  2,
		ParentID:   0, // 会在实际初始化时设置
		Path:       "analytics",
		Name:       "analytics",
		Hidden:     false,
		Component:  "teacher/analytics",
		Title:      "数据分析",
		Icon:       "bar-chart",
		Sort:       3,
		Meta:       systemModel.Meta{KeepAlive: true, DefaultMenu: false, Icon: "bar-chart", Title: "数据分析"},
	}

	// 这里需要将菜单添加到系统中，实际实现时需要调用相应的服务
	// 暂时只定义菜单结构
}
