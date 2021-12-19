# try-go-plugin
本工程是go-plugin框架的一个完整的demo

使用方法
在工程根目录中生成插件可执行文件
```shell
go build -o helloPlugin.exe plugin/plugin.go
```

运行server/server.go


# proto generate
```shell
protoc -I proto/ proto/print.proto --go_out=plugins=grpc:proto/ --go_out=. --go_opt=paths=source_relative

# 分离grpc 的代码到单独文件
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloplugin.proto
```