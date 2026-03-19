package plugin

import (
	"fmt"
	"os/exec"
	"sync/atomic"
)

type PluginStatus int

const (
	Running PluginStatus = iota
	Stopped
	Error
)

// Plugin 插件定义
type Plugin interface {
	Name() string                                                    // 插件名称
	Version() string                                                 // 插件版本
	Run(data map[string]interface{}) (map[string]interface{}, error) // 插件执行方法
	Status() PluginStatus                                            // 插件状态
	SetStatus(PluginStatus)
}

// ExecutionChain 插件执行链
type ExecutionChain struct {
	Plugins []Plugin
}

// PluginManager 插件管理器
type PluginManager struct {
	CurrentChain atomic.Pointer[ExecutionChain] // 存储当前执行链快照
}

type PluginImpl struct {
	PluginName    string
	PluginVersion string
	PStatus       PluginStatus
	ExePath       string
}

func (p *PluginImpl) Name() string {
	return p.PluginName
}

func (p *PluginImpl) Version() string {
	return p.PluginVersion
}

// Run 执行插件
func (p *PluginImpl) Run(data map[string]interface{}) (map[string]interface{}, error) {

	resMap := map[string]interface{}{}
	var err error

	// 参数转换
	var args []string
	for key, val := range data {
		arg := fmt.Sprintf("%s=%s", key, val)
		args = append(args, arg)
	}

	// 创建命令对象并执行
	cmd := exec.Command(p.ExePath, args...)
	output, err := cmd.CombinedOutput()
	outputStr := string(output)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error executing plugin %s: %s", p.ExePath, err.Error()))
	}

	resMap[p.PluginName] = outputStr
	return resMap, err
}

func (p *PluginImpl) Status() PluginStatus {
	return p.PStatus
}

func (p *PluginImpl) SetStatus(status PluginStatus) {
	p.PStatus = status
}
