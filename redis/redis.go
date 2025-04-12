package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func stringOperation(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "大脸猫"

	// 设置键值
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("设置键值失败: %v", err)
	}

	// 获取值
	value1, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("获取键值失败: %v", err)
	}
	fmt.Println("获取的值:", value1)

	// 删除键 (注意方法名是Del不是Delete)
	_, err = client.Del(ctx, key).Result()
	if err != nil {
		log.Fatalf("删除键失败: %v", err)
	}
}

func main() {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 无密码
		DB:       0,  // 默认DB
	})

	// 测试连接
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接Redis: %v", err)
	}

	// 执行字符串操作
	stringOperation(ctx, client)
}
