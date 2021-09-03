FROM golang:1.16.7 as build
WORKDIR /go/src/github.com/zbwang163/
ADD https://api.github.com/repos/zbwang163/ad_info_platform/git/refs/heads/master version.json
RUN git clone -b master https://github.com/zbwang163/ad_info_platform.git \
    && cd ad_info_platform \
    && go mod tidy \
    && echo '-----依赖下载成功-----' \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ad_info_platform . \
    && echo '-----编译成功-----'

FROM alpine:latest as final
RUN apk --no-cache add ca-certificates
WORKDIR /root/app/ad_info_platform
COPY --from=build /go/src/github.com/zbwang163/ad_info_platform/ad_info_platform .
EXPOSE 8080
CMD ["./ad_info_platform"]
