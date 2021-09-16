FROM ubuntu:18.04 AS base

RUN mkdir -p /usr/local/datakit \
    && mkdir -p /usr/local/datakit/externals \
    && mkdir -p /opt/oracle

COPY dist/datakit-linux-amd64/datakit /usr/local/datakit/datakit
COPY dist/datakit-linux-amd64/externals /usr/local/datakit/externals

RUN sed -i 's/\(archive\|security\).ubuntu.com/mirrors.aliyun.com/' /etc/apt/sources.list \
    && apt-get update \
    && apt-get install -y libaio-dev libaio1 unzip wget curl

# download 3rd party libraries
RUN wget -q https://zhuyun-static-files-production.oss-cn-hangzhou.aliyuncs.com/otn_software/instantclient/instantclient-basiclite-linux.x64-19.8.0.0.0dbru.zip \
			-O /usr/local/datakit/externals/instantclient-basiclite-linux.zip \
			&& unzip /usr/local/datakit/externals/instantclient-basiclite-linux.zip -d /opt/oracle

# download data files required by datakit
RUN wget -q -O data.tar.gz https://static.guance.com/datakit/data.tar.gz \
	&& tar -xzf data.tar.gz -C /usr/local/datakit && rm -rf data.tar.gz

CMD ["/usr/local/datakit/datakit", "--docker"]
