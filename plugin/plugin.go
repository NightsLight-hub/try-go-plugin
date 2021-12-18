/*
@Time : 2021/12/18 11:53
@Author : sunxy
@File : plugin
@description:
*/
package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/sxy/try-go-plugin/shared"
)

type PluginService struct{}

func (p PluginService) Hello(name string) (string, error) {
	return "hello " + name, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"PrintPlugin": &shared.GRPCHelloPlugin{Impl: &PluginService{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
