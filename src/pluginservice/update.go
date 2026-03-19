package pluginservice

import (
	"fmt"
	"plugin-based-excutor/src/global"
	"plugin-based-excutor/src/plugin"
)

// UpdateChain 更新当前执行链
func UpdateChain(newPlugins []plugin.Plugin) {
	newChain := &plugin.ExecutionChain{
		Plugins: newPlugins,
	}
	// 对全局 manager 中执行链进行替换
	global.PluginManager.CurrentChain.Store(newChain)
}

// UpdatePluginStatus 更新某插件状态
func UpdatePluginStatus(name string, pluginStatus plugin.PluginStatus) {
	// 获取旧执行链快照
	oldChain := global.PluginManager.CurrentChain.Load()
	if oldChain == nil {
		return
	}
	found := false

	// 更新状态 生成新的执行链
	newPlugins := make([]plugin.Plugin, 0, len(oldChain.Plugins))

	for i := 0; i < len(oldChain.Plugins); i++ {
		p := oldChain.Plugins[i]
		if p.Name() == name {
			fmt.Println("plugin found, new status", pluginStatus)
			// 匹配到后设置状态
			found = true
			p.SetStatus(pluginStatus)
			fmt.Println("plugin status:", p.Status())
		}
		newPlugins = append(newPlugins, p)
	}

	if !found {
		fmt.Println("plugin ", name, " not found")
		return
	}

	// 更新全局插件状态
	UpdateChain(newPlugins)
	fmt.Println(newPlugins)
}
