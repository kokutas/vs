###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-20 11:54:10
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-20 11:55:57
 # @FilePath: /vs/script/volume.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -e
# docker: Error response from daemon: failed to create shim: OCI runtime create failed: container_linux.go:380: starting container process caused: exec: "app": executable file not found in $PATH: unknown.
# fix it
# test -d <文件>：文件存在且是一个目录
volume=$PWD/data
if test -d "$volume"; then
    rm -rf $volume
fi