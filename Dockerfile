FROM alpine:3.3

ENV R11K_VERSION 0.9.2

RUN apk add --update wget && \
    wget -q --no-check-certificate -O /bin/r11k "https://github.com/danielpalstra/r11k/releases/download/${R11K_VERSION}/r11k-linux-amd64" && \
    chmod u+x /bin/r11k && \
    apk del --purge wget && \
    rm -rf /var/cache/apk/*

VOLUME ["/data"]
WORKDIR /data

ENTRYPOINT ["/bin/r11k"]

CMD ["--help"]
