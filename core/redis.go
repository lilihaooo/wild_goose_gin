package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"wild_goose_gin/global"
)

func InitRedis() {
	// 初始化 Redis 连接
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,     // Redis 服务器地址
		Password: global.Config.Redis.Password, // Redis 密码
		DB:       global.Config.Redis.DB,       // 使用的数据库索引
		//ReadTimeout:  20 * time.Second, // 读取超时时间
		//WriteTimeout: 20 * time.Second,
	})
	ctx := context.Background()
	// 确保连接正常
	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.Logrus.Fatal("redis连接失败:", err.Error())
	}
	global.RedisClient = client
}
