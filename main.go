package main

import (
	"fmt"

	"github.com/ArtisanCloud/MediaX/pkg/plugin/core"
)

func main() {
	// 创建 MediaXPlugin 实例
	plugin := &MediaXPlugin{
		PluginName: "MediaXPlugin",
	}

	// 输出插件信息
	fmt.Printf("Plugin name: %s\n", plugin.Name())
	plugin.Publish(core.PublishRequest{
		Title:   "example",
		Content: "Hello, MediaX!",
	})
}
