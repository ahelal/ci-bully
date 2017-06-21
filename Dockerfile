FROM alpine:3.6

ENV CIBULLY_VERSION 0.0.1
ENV CIBULLY_CHECKSUM c2e299cf781b6e79648f9f54b24441b3aba4897a6ed0295b38f3c6ca488bd101

ADD https://github.com/ahelal/ci-bully/releases/download/v${CIBULLY_VERSION}/ci-bully_${CIBULLY_VERSION}_linux_amd64 /usr/bin/cibully
RUN set -ex \
    && echo "${CIBULLY_CHECKSUM}  /usr/bin/cibully" | sha256sum -c \
    && chmod +x /usr/bin/cibully \
    && mkdir /config \
    && apk add --update-cache --no-cache ca-certificates \
    && rm -rf /var/cache/apk/*
VOLUME /config
WORKDIR /config

CMD [ "/bin/sh", "-c", "echo You need to run dockers with cibuly ARGS" ]
