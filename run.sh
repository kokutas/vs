###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-16 00:49:07
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-20 03:52:10
 # @FilePath: /vs/avc/vs/run.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -x
# 停止服务集群，并删除容器和镜像
# docker-compose down --rmi=all
docker-compose down --rmi=local
# docker: Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: exec: "app": executable file not found in $PATH: unknown.
# fix it
# test -d <文件>：文件存在且是一个目录
if test -d "$PWD/data"; then
    rm -rf $PWD/data
fi
# golang:1.18.0
golang=$(docker images|grep golang|grep "1.18.0")
# test -n: 比较字符串的长度是否大于0
# test -z: 比较字符串的长度是否等于0
if test -n "$golang" ; then
    echo $golang
else
   docker pull golang:1.18.0
fi
# alpine:latest
alpine=$(docker images|grep alpine|grep latest)
# test -n: 比较字符串的长度是否大于0
# test -z: 比较字符串的长度是否等于0
if test -n "$alpine" ; then
    echo $alpine
else
   docker pull alpine:latest
fi
# nginx:alpine
nginx=$(docker images|grep nginx|grep alpine)
# test -n: 比较字符串的长度是否大于0
# test -z: 比较字符串的长度是否等于0
if test -n "$nginx" ; then
    echo $nginx
else
   docker pull nginx:alpine
fi
# haproxy:2.5.5
haproxy=$(docker images|grep haproxy|grep "2.5.5")
# test -n: 比较字符串的长度是否大于0
# test -z: 比较字符串的长度是否等于0
if test -n "$haproxy" ; then
    echo $haproxy
else
   docker pull haproxy:2.5.5
fi
docker-compose build --force-rm --no-cache
# docker-compose up -d
# docker ps -a
docker-compose up

# https://unit.nginx.org/configuration/#go
# docker pull nginx/unit:1.26.1-go1.17