###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-20 11:06:25
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-20 12:17:55
 # @FilePath: /vs/script/ip.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -e
sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
&& sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
&& sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
apt update -y && apt install -y tzdata iputils-ping net-tools
ifconfig eth0|grep inet|grep -v inet6|grep -E -o "([0-9]{1,3}[\.]){3}[0-9]{1,3}"|head -n1>>/ip.txt