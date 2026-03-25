package pluginservice

import (
	"fmt"
	"os"
	"plugin-based-excutor/src/common/localconst"
	"plugin-based-excutor/src/plugin"
)

// init 初始化读取最初的插件执行链
func init() {
	fmt.Println("enter init")
	chain := GetExecutionChain()
	UpdateChain(chain)
}

// GetExecutionChain 从路径中获取路径下脚本并排序 生成插件执行链
func GetExecutionChain() []plugin.Plugin {
	fmt.Println("enter GetExecutionChain")

	filePath := localconst.PluginPath
	dirEntries, err := os.ReadDir(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error reading plugin directory: %s", filePath))
		return nil
	}

	var plugins []plugin.Plugin
	for _, dirEntry := range dirEntries {
		fullPath := filePath + "/" + dirEntry.Name()
		fmt.Println(fullPath)
		p := CreateSinglePlugin(fullPath)
		plugins = append(plugins, p)
	}
	return plugins
}

// CreateSinglePlugin 根据单个插件路径 生成单个插件对象
func CreateSinglePlugin(pluginPath string) plugin.Plugin {
	impl := plugin.PluginImpl{
		PluginName:    pluginPath,
		PluginVersion: localconst.DefaultVersion,
		ExePath:       pluginPath,
		PStatus:       plugin.Running,
	}
	return &impl
}
