FROM golang:latest

MAINTAINER <284077318@qq.com> 

COPY . $GOPATH/src/myblog

WORKDIR $GOPATH/src/myblog

RUN go get myblog 

RUN go install -a myblog

# EXPOSE 8080

CMD myblog