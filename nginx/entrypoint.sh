###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-17 23:54:39
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-18 02:42:05
 # @FilePath: /vs/nginx/entrypoint.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
###
#!/bin/sh
/usr/sbin/keepalived -n -l -D -f /etc/keepalived/keepalived.conf --dont-fork --log-console &

nginx -g "daemon off;"