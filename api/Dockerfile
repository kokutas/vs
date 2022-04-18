# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.18.0 AS builder
# 启用go module
ENV GO111MODULE=on \
GOPROXY=https://goproxy.io,https://gonexus.dev,https://athens.azurefd.net,https://mirrors.aliyun.com/goproxy,https://proxy.golang.com.cn,https://goproxy.cn
# 指定工作目录
WORKDIR /workspace
COPY . .
RUN go mod tidy \
&& go mod vendor \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o api ./cmd/main.go
# RUN GOOS=linux GOARCH=amd64 go build --ldflags "-extldflags -static" -v -o api ./cmd/main.go

# FROM --platform=$BUILDPLATFORM alpine:latest
FROM alpine:latest
LABEL author="kokutas" email="2150748714@qq.com" api="v1.0.0" description="restful api v1.0.0 by golang v1.8.0 docker image"
USER root
ENV TIME_ZONE=Asia/Shanghai \
GIN_MODE=release \
PORT=8080 \
NAMESPACE="api gateway" \
API_LOG_DIR=/var/api/log \
API_LOG_PREFIX="api" \
API_LOG_LEVEL="debug" \
API_LOG_DEVELOPMENT=true \
API_LOG_DEBUG_FILE_SUFFIX="debug.log" \
API_LOG_WARN_FILE_SUFFIX="warn.log" \
API_LOG_ERROR_FILE_SUFFIX="error.log" \
API_LOG_INFO_FILE_SUFFIX="info.log" \
API_LOG_MAX_AGE=30 \
API_LOG_MAX_BACKUPS=100 \
API_LOG_MAX_SIZE=1 \
PATH=$PATH:/usr/local/bin

RUN mkdir -p ${API_LOG_DIR} \
&& sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk add ca-certificates \
&& rm -rf /var/cache/apk/* \
&& apk add --no-cache bash tzdata \
&& cp /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime \
&& echo ${TIME_ZONE} > /etc/timezone
WORKDIR /usr/local/bin
# 将上一个阶段编译好的二进制文件app复制进来
COPY --from=builder /workspace/api ./
VOLUME ${API_LOG_DIR}
# 开放端口
EXPOSE ${PORT}
ENTRYPOINT ["api"]