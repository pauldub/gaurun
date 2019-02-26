FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ADD bin/linux/amd64/gaurun-0.10.0 /

USER nobody:nobody

CMD ["/gaurun"]

