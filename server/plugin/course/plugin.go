package course

import (
	"github.com/flipped-aurora/gin-vue-admin/server/interfaces"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/course/initialize"
	"github.com/gin-gonic/gin"
)

// Plugin 课程插件
type Plugin struct{}

// Register 注册插件
func (p *Plugin) Register(router *gin.RouterGroup) {
	// 初始化路由
	initialize.InitRouter(router)
}

// RouterPath 返回插件路由路径
func (p *Plugin) RouterPath() string {
	return "/course"
}

// Name 返回插件名称
func (p *Plugin) Name() string {
	return "course"
}

// Init 初始化插件
func (p *Plugin) Init() {
	// 初始化数据库
	initialize.InitDB(nil) // 实际使用时会传入数据库连接
	
	// 初始化菜单
	initialize.InitMenu()
}

// Plugin 插件实例
var Plugin = new(Plugin)

func init() {
	// 注册插件到系统
	interfaces.Register(Plugin)
}
