package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "192.168.43.124:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}