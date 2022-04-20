###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-20 11:44:13
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-20 11:48:11
 # @FilePath: /vs/script/images.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -e
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
# nginx:latest
nginx=$(docker images|grep nginx|grep latest)
# test -n: 比较字符串的长度是否大于0
# test -z: 比较字符串的长度是否等于0
if test -n "$nginx" ; then
    echo $nginx
else
   docker pull nginx:latest
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
