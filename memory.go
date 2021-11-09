package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func ConnectToRedis(data []byte) {
	fmt.Println("Config to open redis db")

	// var client redis.Client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       7,
	})

	err := client.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(pong, err)
	// return client

	err = client.Set(time.Now().Format("2006-01-02"), data, 0).Err()
	if err != nil {
		fmt.Print("Something happened when saving")
		log.Fatal(err)
	}
	fmt.Print("Successfully saved to Redis")

}
