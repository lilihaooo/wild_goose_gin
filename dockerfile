# MySQL 阶段
FROM mysql:latest AS mysql

# 设置 MySQL 环境变量
ENV MYSQL_ROOT_PASSWORD=a111111
ENV MYSQL_DATABASE=wild_goose_gin:1111
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=a111111
ENV LANG=C.UTF-8

# 设置工作目录
WORKDIR /mysql

# 复制自定义的 my.cnf 文件到 /etc/mysql/my.cnf
COPY mysql/conf/my.cnf /etc/mysql/my.cnf

# 暴露 MySQL 默认端口
EXPOSE 3306

# 设置启动 MySQL 的命令
CMD ["mysqld"]

# Redis 阶段
FROM redis:latest AS redis

# 暴露 Redis 默认端口
EXPOSE 6379

# 设置工作目录
WORKDIR /redis

# 复制自定义的 redis.conf 文件到 /etc/redis/redis.conf
COPY redis/conf/redis.conf /etc/redis/redis.conf

# 设置启动 Redis 的命令
CMD ["redis-server", "/etc/redis/redis.conf"]

# 构建阶段
FROM golang:latest AS builder

# 设置工作目录
WORKDIR /go/src/wild_goose_gin

# 复制 go.mod 和 go.sum 文件并下载依赖项
COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制整个项目并构建 Go 项目
COPY . .
RUN go build -o main .

# 从 MySQL 阶段复制初始化脚本
COPY --from=mysql /etc/mysql/my.cnf /etc/mysql/my.cnf

# 从 Redis 阶段复制 Redis 配置文件
COPY --from=redis /etc/redis/redis.conf /etc/redis/redis.conf

# 设置 MySQL 和 Redis 地址的环境变量
ENV MYSQL_ADDR=localhost:3306
ENV REDIS_ADDR=localhost:6379

# 暴露应用程序的端口
EXPOSE 8080

# 设置运行 Go 可执行文件的命令
CMD ["./main"]
