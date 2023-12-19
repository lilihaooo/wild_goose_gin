# 构建阶段
FROM golang:latest AS builder

# 设置工作目录
WORKDIR /app

# 复制整个项目并构建 Go 项目
COPY . .
RUN go build -o main .

# 暴露应用程序的端口
EXPOSE 8888

# 设置运行 Go 可执行文件的命令
CMD ["./main"]
