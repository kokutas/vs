# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.18.0 AS builder
# 启用go module
ENV GO111MODULE=on \
GOPROXY=https://goproxy.io,https://gonexus.dev,https://athens.azurefd.net,https://mirrors.aliyun.com/goproxy,https://proxy.golang.com.cn,https://goproxy.cn
# 指定工作目录
WORKDIR /workspace
COPY . .
RUN go mod tidy \
&& go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o app ./cmd/main.go
# FROM --platform=$BUILDPLATFORM alpine:latest
FROM alpine:latest
LABEL author="kokutas" email="2150748714@qq.com" logs="v1.0.0" description="area v1.0.0 by golang v1.8.0 docker image"
USER root

ENV TIME_ZONE=Asia/Shanghai \
# 日志文件路径,自动创建目录
APP_LOG_DIR=/var/app/log \
# 日志文件前缀
APP_LOG_PREFIX="app" \
# 日志级别：debug/info/error/warn
APP_LOG_LEVEL="debug" \
# 是否为开发者模式
APP_LOG_DEVELOPMENT=true \
# debug日志文件后缀
APP_LOG_DEBUG_FILE_SUFFIX="debug.log" \
# warn日志文件后缀
APP_LOG_WARN_FILE_SUFFIX="warn.log" \
# error日志文件后缀
APP_LOG_ERROR_FILE_SUFFIX="error.log" \
# info日志文件后缀
APP_LOG_INFO_FILE_SUFFIX="info.log" \
# 日志文件保存的最大天数
APP_LOG_MAX_AGE=30 \
# 日志文件最多存在多少个切片文件
APP_LOG_MAX_BACKUPS=100 \
# 日志文件大小（M）
APP_LOG_MAX_SIZE=1 \
PATH=$PATH:/usr/local/bin

RUN mkdir -p ${APP_DATA_DIR} ${APP_LOG_DIR} \
&& sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk add ca-certificates \
&& rm -rf /var/cache/apk/* \
&& apk add --no-cache bash tzdata \
&& cp /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime \
&& echo ${TIME_ZONE} > /etc/timezone
WORKDIR /usr/local/bin
# 将上一个阶段编译好的二进制文件app复制进来
COPY --from=builder /workspace/app ./
VOLUME /var/app/data, /var/app/data
ENTRYPOINT ["app"]




FROM golang

RUN mkdir -p /go/src/vs

ADD . /go/src/vs

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run vs/area/cmd  -watch vs/area
