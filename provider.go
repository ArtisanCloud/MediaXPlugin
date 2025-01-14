package main

import (
	"fmt"

	"github.com/ArtisanCloud/MediaX/pkg/plugin/core"
)

type MediaXPlugin struct {
	PluginName string
}

func (p *MediaXPlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "MediaXPlugin"
	return nil
}

func (p *MediaXPlugin) Name() string {
	return p.PluginName
}

func (p *MediaXPlugin) Publish(core.PublishRequest, ...interface{}) (core.PublishResult, error) {

	fmt.Println("Publishing MediaX Plugin")
	return core.PublishResult{
		Status:  "success",
		Message: "MediaX Plugin Published Successfully",
	}, nil
}

// 导出插件实例
var Provider core.Provider = &MediaXPlugin{
	PluginName: "MediaXPlugin",
}
