FROM nginx:alpine
LABEL author="kokutas" email="2150748714@qq.com" nginx="alpine" description="add time zone based on nginx:alpine"
USER root
COPY entrypoint.sh /entrypoint.sh
ENV TIME_ZONE=Asia/Shanghai 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk upgrade \
&& apk add --no-cache tzdata bash curl ipvsadm iproute2 openrc keepalived \
&& ln -snf /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime && echo ${TIME_ZONE} > /etc/timezone \
&& rm -f /var/cache/apk/* /tmp/* \
&& chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]