FROM haproxy:2.5.5
LABEL author="kokutas" email="2150748714@qq.com" haproxy="2.5.5" description="add time zone based on haproxy:2.5.5"
USER root
USER root
ENV TIME_ZONE=Asia/Shanghai
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
&& apt-get update \
&& ln -snf /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime && echo ${TIME_ZONE} > /etc/timezone \
&& apt-get install tzdata \
&& apt-get clean \
&& apt-get autoclean \
&& rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
RUN dpkg-reconfigure -f noninteractive tzdata