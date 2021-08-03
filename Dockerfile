FROM golang

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /scorelist
ADD . /app
ENTRYPOINT ["./run.sh"]
EXPOSE 9000
ENTRYPOINT ["run"]

FROM golang:1.9 as builder
# define stage name as builder￼
RUN mkdir -p /go/src/test
WORKDIR /go/src/test
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/test/app .
CMD ["./app"]