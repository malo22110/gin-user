FROM golang:alpine

ADD . $GOPATH/src/gin-user-service

RUN set -ex \
  && apk add --update git \
  && go get -u github.com/gin-gonic/gin \
  && go get -u github.com/jinzhu/gorm \
  && cd $GOPATH/src/gin-user-service \
  && ls . \
  && go install

VOLUME /data

CMD ["gin-user-service"]