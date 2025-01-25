package main

import (
	"fmt"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
)

type MediaXPlugin struct {
	PluginName string
}

func NewMediaXPlugin() MediaXPlugin {
	return MediaXPlugin{
		PluginName: "PluginMediaX",
	}
}

func (p *MediaXPlugin) Initialize(config map[string]interface{}) error {
	p.PluginName = "PluginMediaX"
	return nil
}

func (p *MediaXPlugin) Name() string {
	return p.PluginName
}

func (p *MediaXPlugin) Publish(*contract.PublishRequest, ...interface{}) (*contract.PublishResult, error) {

	fmt.Println("Publishing MediaX Plugin")
	return &contract.PublishResult{
		Status:  "success",
		Message: "MediaX Plugin Published Successfully",
	}, nil
}

var PluginMediaX MediaXPlugin = NewMediaXPlugin()
