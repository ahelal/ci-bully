FROM alpine:3.6

ENV CIBULLY_VERSION 0.0.1
ENV CIBULLY_CHECKSUM 176bdfcdb17e757b9baec7f4984baf6c58127f85b43067f0482e307b08740336

ADD https://github.com/ahelal/ci-bully/releases/download/v${CIBULLY_VERSION}/cibully_${CIBULLY_VERSION}_linux_amd64 /usr/bin/cibully
RUN set -ex \
    && echo "${CIBULLY_CHECKSUM}  /usr/bin/cibully" | sha256sum -c \
    && chmod +x /usr/bin/cibully \
    && mkdir /config \
    && apk add --update-cache --no-cache ca-certificates \
    && rm -rf /var/cache/apk/*
VOLUME /config
WORKDIR /config

CMD [ "/bin/sh", "-c", "echo You need to run dockers with cibuly ARGS" ]
