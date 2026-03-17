package main

import (
	_ "plugin-based-excutor/src/pluginservice"
	"plugin-based-excutor/src/pluginweb"
)

func main() {
	pluginweb.PluginStarterMain()
}
