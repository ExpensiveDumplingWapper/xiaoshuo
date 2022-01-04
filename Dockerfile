FROM uhub.service.ucloud.cn/bluecity/golang:1.17.0-alpine as builder
LABEL maintainer="lizhongzheng@blued.com"

# 基于go moudles, 我们把源码放到root目录下进行编译
WORKDIR /root

# add source code to WORKDIR (do not ignore vendor)
COPY . .

# build -ldflags="-s -w" 移除debug信息, distroless/base 镜像提供了glibc的支持 所以移除关闭cgo的操作 CGO_ENABLED=0
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -o main app/main.go

# 构建服务运行容器
FROM uhub.service.ucloud.cn/bluecity/alpine:3.14

ENV APP_PORT 80
ENV IPDB_PATH /app/ipdb
# copy the binary from builder
COPY --from=builder /root .
EXPOSE $APP_PORT
# run the binary
CMD ["./main"]