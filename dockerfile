# 导入基础镜像golang:alpine
FROM golang:alpine AS builder

# 设置环境变量
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

# 创建并移动到工作目录（可自定义路径）
WORKDIR /app

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件,文件名为 WebApp
RUN go build -o main .

# 声明服务端口
EXPOSE 8888

# 启动容器时运行的命令
CMD ["./main"]

