FROM golang:latest

MAINTAINER <284077318@qq.com> 

COPY . $GOPATH/src/github.com/jicg/myblog

WORKDIR $GOPATH/src/github.com/jicg/myblog

RUN go get github.com/jicg/myblog 

RUN go install -a github.com/jicg/myblog

# EXPOSE 8080

CMD myblog