<!--
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-16 02:35:14
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-18 01:33:10
 * @FilePath: /vs/README_ZH.md
 * @Description: TODO
 * Copyright (c) 2022 by kokutas, All Rights Reserved. 
-->
# 高可用方案
> keepalived + nginx
> 
|vip|服务器|nginx|
|:--:|:--:|:--:|
|10.0.0.2|10.0.0.3|master|
|10.0.0.2|10.0.0.4|slave|

# 项目
## 前提
```shell
go mod init github.com/kokutas/vs
go get github.com/GrantZheng/kit@latest
go install github.com/GrantZheng/kit@latest
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc@latest
go get google.golang.org/protobuf@latest
```
## 地区（area）
- 创建项目
```shell
kit new service area --module github.com/kokutas/vs/area
```