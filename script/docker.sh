###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-20 11:07:04
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-20 11:48:56
 # @FilePath: /vs/script/docker.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -e
docker_version=$(docker info|grep -i -E ".*server.*version.*"|grep -E -o "[0-9]{1,3}"|head -n1)
if [ $docker_version -ge 20 ]; then
    echo $(docker info|grep -i -E ".*server.*version.*")
else 
    echo "docker engine server version must greater or equal to 20.x.x"
    exit 1
fi
docker_compose_version=$(docker-compose version|grep -E -o "[0-9]{1,3}"|head -n1)
if [ $docker_compose_version -ge 2 ]; then
    echo $(docker-compose version)
else 
    echo "docker-compose version must greater or equal to v2.x.x"
    exit 1
fi