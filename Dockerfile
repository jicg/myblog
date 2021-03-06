ARG GO_VERSION=1.11
FROM golang:${GO_VERSION}-alpine AS builder
MAINTAINER <284077318@qq.com>
COPY . $GOPATH/src/github.com/jicg/myblog
WORKDIR $GOPATH/src/github.com/jicg/myblog
RUN apk add --no-cache ca-certificates git
RUN go get github.com/jicg/myblog
RUN CGO_ENABLED=0 go install -a github.com/jicg/myblog
# EXPOSE 8080

# CMD myblog
FROM scratch AS final
MAINTAINER <284077318@qq.com>
COPY --from=builder /go/bin/myblog /app/myblog
#/usr/bin/myblog
COPY --from=builder /go/src/github.com/jicg/myblog/views /app/views
COPY --from=builder /go/src/github.com/jicg/myblog/static /app/static
COPY --from=builder /go/src/github.com/jicg/myblog/sys.ini /app/sys.ini
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
VOLUME /app/data
VOLUME /app/conf
VOLUME /app/cert-cache

EXPOSE 80
#EXPOSE 443
WORKDIR /app
ENTRYPOINT ["/app/myblog"]