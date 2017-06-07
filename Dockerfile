FROM alpine:3.6

ENV CIBULLY_VERSION 0.0.1-rc.1
ENV CIBULLY_CHECKSUM ccecd496ab981ea2db56bdf5737a96e2adcabdfcf4bea53a002fdaa893b39311

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
