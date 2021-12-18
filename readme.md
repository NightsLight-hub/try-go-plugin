# proto generate
```shell
protoc -I proto/ proto/print.proto --go_out=plugins=grpc:proto/ --go_out=. --go_opt=paths=source_relative

# 分离grpc 的代码到单独文件
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloplugin.proto
```