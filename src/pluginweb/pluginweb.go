package pluginweb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"plugin-based-excutor/src/common/localconst"
	"plugin-based-excutor/src/plugin"
	"plugin-based-excutor/src/pluginservice"
	"strconv"
)

// PluginStarterMain PluginStarter 启动入口
func PluginStarterMain() {
	http.HandleFunc("/execute", PluginExecuteHandler)
	http.HandleFunc("/update", PluginUpdateHandler)
	http.HandleFunc("/getstatus", PluginGetHandler)

	addr := ":" + localconst.ListenPort
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error starting http server ", err)
		return
	}
}

// PluginExecuteHandler 执行插件 对外接口 handler 实现
func PluginExecuteHandler(w http.ResponseWriter, r *http.Request) {
	// 处理 post 请求
	bodyBytes, err := io.ReadAll(r.Body)
	response := map[string]interface{}{}

	if err != nil {
		fmt.Println("Error reading body ", err)
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	inputData := map[string]interface{}{}
	err = json.Unmarshal(bodyBytes, &inputData)
	if err != nil {
		fmt.Println("Error unmarshalling body ", err)
		http.Error(w, "Error unmarshalling body", http.StatusInternalServerError)
		return
	}

	// 执行插件逻辑
	res, err := pluginservice.HandleRequest(inputData)
	if err != nil {
		fmt.Println("Error handling request", err)
		response["error"] = err.Error()
	}
	response["data"] = res

	// 响应数据
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Error encoding response", err)
		return
	}
}

// PluginUpdateHandler 更新插件 对外接口 handler 实现
func PluginUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// 处理 post 请求
	bodyBytes, err := io.ReadAll(r.Body)
	response := map[string]interface{}{}

	if err != nil {
		fmt.Println("Error reading body ", err)
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	inputData := map[string]interface{}{}
	err = json.Unmarshal(bodyBytes, &inputData)
	if err != nil {
		fmt.Println("Error unmarshalling body ", err)
		http.Error(w, "Error unmarshalling body", http.StatusInternalServerError)
		return
	}

	pName := inputData[localconst.UpdatePluginNameParam]
	pStatus := inputData[localconst.UpdatePluginStatusParam]
	pStatusVal, err := PStatusStrToVal(pStatus.(string))
	if err != nil {
		fmt.Println("Error converting status to val", err)
		http.Error(w, "Error converting status to val", http.StatusInternalServerError)
		return
	}

	pluginservice.UpdatePluginStatus(pName.(string), pStatusVal)
	// 响应数据
	w.WriteHeader(http.StatusOK)
	response["result"] = true
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Error encoding response", err)
		return
	}
}

// PluginGetHandler 获取 plugin 当前状态的接口 对外接口 handler 实现
func PluginGetHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pluginName := queryParams.Get(localconst.QueryPluginNameParam)

	p, err := pluginservice.GetPluginStatusByName(pluginName)
	if err != nil {
		fmt.Println("Error getting plugin", err)
		return
	}

	msg := PluginMsg{
		Name:    p.Name(),
		Version: p.Version(),
		Status:  p.Status(),
	}

	// 响应数据
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(msg)
	if err != nil {
		fmt.Println("Error encoding response", err)
		return
	}

}

// PStatusStrToVal 转换 string val 到 PluginStatus
func PStatusStrToVal(val string) (plugin.PluginStatus, error) {
	rawNum, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Error converting atoi", err)
		return plugin.Error, err
	}
	status := plugin.PluginStatus(rawNum)
	if status >= plugin.Running && status <= plugin.Error {
		return status, nil
	}
	return plugin.Error, errors.New("plugin status out of range")
}
