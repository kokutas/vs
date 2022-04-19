FROM haproxy:2.5.5
LABEL author="kokutas" email="2150748714@qq.com" haproxy="2.5.5" description="add time zone based on haproxy:2.5.5"
USER root
USER root
ENV TZ=Asia/Shanghai
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
&& sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
&& sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
&& apt-get update \
&& ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime && echo ${TZ} > /etc/timezone \
&& apt-get install tzdata \
&& apt-get clean \
&& apt-get autoclean \
&& rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
RUN dpkg-reconfigure -f noninteractive tzdata