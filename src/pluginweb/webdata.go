package pluginweb

import "plugin-based-excutor/src/plugin"

type Response struct {
	Data map[string]string `json:"data"`
}

type PluginMsg struct {
	Name    string              `json:"name"`
	Status  plugin.PluginStatus `json:"status"`
	Version string              `json:"version"`
}
