/*
@Time : 2021/12/18 11:52
@Author : sunxy
@File : plugin_interface
@description:
*/
package shared

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/sxy/try-go-plugin/proto"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

type IHelloService interface {
	Hello(name string) (string, error)
}

// GRPCHelloPlugin implement plugin.GRPCPlugin
type GRPCHelloPlugin struct {
	plugin.Plugin
	Impl IHelloService
}

func (p GRPCHelloPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	proto.RegisterHelloPluginServer(server, GPRCHelloPluginServerWrapper{impl: p.Impl})
	return nil
}

func (p GRPCHelloPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return GRPCHelloPluginClientWrapper{client: proto.NewHelloPluginClient(conn)}, nil
}

type GPRCHelloPluginServerWrapper struct {
	impl IHelloService
	proto.UnimplementedHelloPluginServer
}

func (_this GPRCHelloPluginServerWrapper) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	r, _ := _this.impl.Hello(request.Name)
	return &proto.Response{
		Result: r,
	}, nil
}

// GRPCHelloPluginClientWrapper 作为server 调用插件接口的包装器，
type GRPCHelloPluginClientWrapper struct {
	client proto.HelloPluginClient
}

func (_this GRPCHelloPluginClientWrapper) Hello(name string) (string, error) {
	in := proto.Request{Name: name}
	resp, err := _this.client.Hello(context.Background(), &in)
	if err != nil {
		return "", err
	} else {
		return resp.Result, nil
	}
}
