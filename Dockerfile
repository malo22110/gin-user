FROM golang:alpine

ADD . $GOPATH/src/gin-user-service

RUN set -ex \
  && apk add --update git \
  && cd $GOPATH/src/gin-user-service \
  && ls . \
  && go install

VOLUME /data

CMD ["gin-user-service"]