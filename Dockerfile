FROM golang:latest
MAINTAINER <284077318@qq.com>
COPY . $GOPATH/src/github.com/jicg/myblog
WORKDIR $GOPATH/src/github.com/jicg/myblog
RUN apk add --no-cache ca-certificates git
RUN go get github.com/jicg/myblog
RUN CGO_ENABLED=0 go install -a github.com/jicg/myblog
# EXPOSE 8080

# CMD myblog
FROM busybox
MAINTAINER <284077318@qq.com>
COPY --from=0 /go/bin/myblog /usr/bin/myblog
#/usr/bin/myblog
COPY --from=0 /go/src/github.com/jicg/myblog/views /app/views
COPY --from=0 /go/src/github.com/jicg/myblog/static /app/static
COPY --from=0 /go/src/github.com/jicg/myblog/sys.ini /app/sys.ini
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
VOLUME /app/data
VOLUME /app/conf
VOLUME /app/cert-cache

EXPOSE 8080
EXPOSE 8443
WORKDIR /app
CMD /usr/bin/myblog