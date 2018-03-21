#test
FROM alpine:edge
MAINTAINER 吕钊华

ADD main /appdir/

RUN apk update && apk upgrade \
  && apk add ca-certificates \
  && rm -rf /var/cache/apk/*

ENTRYPOINT ["./appdir/main"]