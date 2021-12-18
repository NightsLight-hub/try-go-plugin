/*
@Time : 2021/12/18 11:53
@Author : sunxy
@File : server
@description:
*/
package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/sxy/try-go-plugin/shared"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.SetOutput(os.Stdout)
	pluginClientConfig := &plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Cmd:              exec.Command("./helloPlugin.exe"),
		Plugins:          map[string]plugin.Plugin{"main": &shared.GRPCHelloPlugin{}},
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	}

	client := plugin.NewClient(pluginClientConfig)
	pluginClientConfig.Reattach = client.ReattachConfig()
	protocol, err := client.Client()
	if err != nil {
		log.Fatalln(err)
	}
	raw, err := protocol.Dispense("main")
	if err != nil {
		log.Fatalln(err)
	}
	service := raw.(shared.IHelloService)
	res, err := service.Hello("sxy")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
