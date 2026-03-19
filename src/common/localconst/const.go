package localconst

const (
	ListenPort           = "8123"                        // 插件执行器对外 web 服务端口
	PluginPath           = "/app/plugin-excutor/plugins" // 插件路径
	DefaultPluginTimeout = 30                            // 插件执行超时时间 单位：s
	DefaultVersion       = "0.0.1"

	QueryPluginNameParam    = "pluginname"   // 查询 plugin name 字段
	UpdatePluginNameParam   = "pluginname"   // 更新 plugin 状态 插件名称字段
	UpdatePluginStatusParam = "pluginstatus" // 更新 plugin 状态 插件状态字段
)
