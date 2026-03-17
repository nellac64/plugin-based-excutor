package pluginservice

import (
	"context"
	"errors"
	"fmt"
	"plugin-based-excutor/src/common/localconst"
	"plugin-based-excutor/src/global"
	"plugin-based-excutor/src/plugin"
	"time"
)

// HandleRequest 处理请求 依次执行插件
func HandleRequest(inputData map[string]interface{}) (map[string]interface{}, error) {

	resData := map[string]interface{}{}
	singleRes := map[string]interface{}{}
	var err error
	// 加载当前执行链的快照
	chain := global.PluginManager.CurrentChain.Load()
	if chain == nil {
		return resData, errors.New("plugin chain is nil")
	}

	execCtx, cancel := context.WithTimeout(context.Background(), localconst.DefaultPluginTimeout*time.Second)
	defer cancel()

	// 根据执行链 依次执行插件
	for _, p := range chain.Plugins {
		pluginName := p.Name()
		select {
		case <-execCtx.Done():
			return resData, fmt.Errorf("plugin timeout, name: %s", pluginName)
		default:
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("plugin panic: %v, name: %s", r, pluginName)
					}
				}()
				if p.Status() != plugin.Running {
					fmt.Printf("plugin status is %s, name: %s, do not need running", p.Status(), pluginName)
					return
				}
				singleRes, err = p.Run(inputData)
				if err != nil {
					// 单个插件执行报错 打印日志
					fmt.Println(fmt.Sprintf("plugin execute err: %v, name: %s", err, pluginName))
				}
				resData = MapCombine(resData, singleRes)
			}()
		}
	}
	return resData, err
}

// MapCombine map 合并
func MapCombine(baseMap map[string]interface{}, addMap map[string]interface{}) map[string]interface{} {
	for key, val := range addMap {
		baseMap[key] = val
	}
	return baseMap
}
