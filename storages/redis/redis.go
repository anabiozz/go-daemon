package redis

import (
	"fmt"
	"os"

	"log"

	"github.com/go-redis/redis"
)

// Connection (redis connect)
func Connection() *redis.Client {
	redisAddr := os.Getenv("REDIS_PORT_6379_TCP_ADDR")
	var client *redis.Client
	if len(redisAddr) > 0 {
		client = connect(redisAddr)
	} else {
		client = connect("localhost")
	}
	return client
}

func connect(redisAddr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalln("redis error: ", err)
	}
	fmt.Println("redis connection", pong)
	return client
}

// Save (save object to redis)
func Save(key string, value interface{}, client *redis.Client) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("redis error: ", err)
	}

	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println("redis error: ", err)
	}
	fmt.Println("key", val)
}
