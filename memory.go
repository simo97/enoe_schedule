package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectToRedis() {
	fmt.Println("Config to open redis db")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       7,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
