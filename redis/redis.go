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

}

func hashtable(ctx context.Context, client *redis.Client) {
	err := client.HSet(ctx, "学生2", "Name", "大脸猫", "Age", 18, "Height", 173.5).Err()
	if err != nil {
		log.Fatalf("设置哈希表失败: %v", err)
	}
	name, err := client.HGet(ctx, "学生2", "Name").Result()
	if err != nil {
		fmt.Printf("获取Name失败: %v\n", err)
	} else {
		fmt.Println("获取的Name值:", name)
	}

	age, err := client.HGet(ctx, "学生2", "Age").Result()
	if err != nil {
		fmt.Printf("获取Age失败: %v\n", err)
	} else {
		fmt.Println("获取的Age值:", age)
	}

	height, err := client.HGet(ctx, "学生2", "Height").Result()
	if err != nil {
		fmt.Printf("获取Height失败: %v\n", err)
	} else {
		fmt.Println("获取的Height值:", height)
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
	hashtable(ctx, client)
}
