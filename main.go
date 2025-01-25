package main

import (
	"fmt"
	plugin2 "github.com/ArtisanCloud/MediaXCore/pkg/plugin"
	"github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"
	"plugin"
)

func main() {
	// 加载插件
	p, err := plugin.Open("./plugin/plugin.so")
	if err != nil {
		panic(err)
	}

	mediaXPlugin, err := plugin2.LookUpSymbol[contract.ProviderInterface](p, "PluginMediaX")
	if err != nil {
		panic(err)
	}
	fmt.Printf("plugin loaded name :%s", (*mediaXPlugin).Name())

}
