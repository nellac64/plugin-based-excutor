package pluginservice

import (
	"errors"
	"plugin-based-excutor/src/global"
	"plugin-based-excutor/src/plugin"
)

// GetPluginStatusByName 根据 plugin 名称 获取 plugin 状态
func GetPluginStatusByName(pluginName string) (status plugin.Plugin, err error) {
	// 获取当前只读副本
	var result plugin.Plugin
	isFound := false

	nowChain := global.PluginManager.CurrentChain.Load()
	for i := 0; i < len(nowChain.Plugins); i++ {
		if nowChain.Plugins[i].Name() == pluginName {
			result = nowChain.Plugins[i]
			isFound = true
			break
		}
	}
	if !isFound {
		return nil, errors.New("plugin not found")
	}
	return result, nil
}
