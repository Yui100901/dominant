FROM golang:1.23.0-alpine AS builder

# 设置工作目录为 /app
# 所有后续操作都会在这个目录下进行
WORKDIR /app

# 将当前项目目录的所有文件拷贝到容器的 /app 目录中
COPY . .

# 设置 Go 模块代理为 https://goproxy.cn（在中国加速模块下载），并下载项目的依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

# 编译 Go 项目，生成可执行文件 simple-web-app
RUN go build -o dominant

# 使用一个更小的基础镜像（Alpine）来运行应用程序
# Alpine 是一个极简的 Linux 发行版，适合部署阶段
FROM alpine:latest

# 安装 tzdata 包，确保支持时区的配置
RUN apk add --no-cache tzdata

# 设置工作目录为 /app
WORKDIR /app

# 从编译阶段的镜像中拷贝编译后的二进制文件到运行镜像中
COPY --from=builder /app/dominant /app/dominant

# 暴露容器的 8080 端口，用于外部访问
EXPOSE 28888

# 设置容器启动时运行的命令
# 这里是运行编译好的可执行文件 simple-web-app
CMD ["/app/dominant"]