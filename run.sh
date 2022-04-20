###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-20 11:47:17
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-20 13:08:08
 # @FilePath: /vs/run.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -x
# 停止服务集群，并删除容器和镜像
# docker-compose down --rmi=all
# 停止服务集群，并删除容器和当前本地创建的镜像
docker-compose down --rmi=local
docker pull golang:1.18.0
docker pull alpine:latest
docker pull nginx:latest
docker pull haproxy:2.5.5

# test -e <文件>：文件是否存在
dockersh=$PWD/script/docker.sh
if test -e "$dockersh"; then
    sh $dockersh
else
    echo "Unverified docker and docker-compose versions."
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
volumesh=$PWD/script/volume.sh
if test -e "$volumesh"; then
    sh $volumesh
else
    echo "No historical data volume mount and clear."
fi
haproxy_cfg=$PWD/config/haproxy/haproxy.cfg
if test -e "$haproxy_cfg"; then
    sed -e '/server master/d' $haproxy_cfg > tmp.cfg
    sed -e '/server slave/d' tmp.cfg > $haproxy_cfg
    rm -f tmp.cfg
    # 获取当前network_mode:host模式的IP地址
    hostipsh=$PWD/script/ip.sh
    if test -e "$hostipsh"; then
        touch ip.txt
        docker run -it --rm --privileged --user 0 --network host -v $hostipsh:/ip.sh -v $PWD/ip.txt:/ip.txt --name haproxy haproxy:2.5.5 /bin/sh /ip.sh
        cat $PWD/ip.txt|head -n1|xargs -I ip echo "    server master ip:8081 check inter 2000 rise 2 fall 5">>$haproxy_cfg
        cat $PWD/ip.txt|head -n1|xargs -I ip echo "    server slave ip:8082 check inter 2000 rise 2 fall 5">>$haproxy_cfg
        rm -f $PWD/ip.txt
    else
        echo "IP address not found for network_mode:host mode."
    fi
fi
docker-compose build --force-rm --no-cache
# docker-compose up -d
# docker ps -a
docker-compose up

# https://unit.nginx.org/configuration/#go
# docker pull nginx/unit:1.26.1-go1.17
