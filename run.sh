###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-16 00:49:07
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-18 02:25:46
 # @FilePath: /vs/run.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -x
# 停止服务集群，并删除容器和镜像
# docker compose down --rmi=all
docker compose down 
rm -rf data
docker network rm frontend
docker compose up