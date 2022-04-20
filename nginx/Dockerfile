FROM nginx:alpine
LABEL author="kokutas" email="2150748714@qq.com" nginx="alpine" description="add time zone and keepalived based on nginx:alpine"
USER root
COPY entrypoint.sh /entrypoint.sh
ENV TZ=Asia/Shanghai 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk upgrade \
&& apk add --no-cache tzdata bash curl ipvsadm iproute2 openrc keepalived \
&& ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime && echo ${TZ} > /etc/timezone \
&& rm -f /var/cache/apk/* /tmp/* \
&& chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]