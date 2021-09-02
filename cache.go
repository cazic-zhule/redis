package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

const UNSET = "UNSET"

type RedisConfig struct {
	Client *redis.Client
}

func (r *RedisConfig) Exists(key string) bool {
	v := r.Get(key)
	return v != UNSET
}

func (r *RedisConfig) Get(key string) string {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return UNSET
	} else if err != nil {
		fmt.Println("Error getting from redis with key, ", key, " in redis ", err)
		return ""
	}
	return val
}

func (r *RedisConfig) Set(key string, val string, ttl time.Duration) {
	err := r.Client.Set(ctx, key, val, ttl).Err()
	if err != nil {
		fmt.Println("Error setting key, ", key, ", val, ", val, " in redis ", err)
	}
}

/*
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "abc123", // no password set
		DB:       0,        // use default DBredis
	})

	c := &RedisConfig{client: rdb}

	c.Set("fooo", "fffff", time.Second*2)
	time.Sleep(4 * time.Second)
	v := c.Get("fooo")
	fmt.Println("key", v)

	test := c.Exists("fooo")
	fmt.Println("exists, fooo", test)

	c.Set("hello", "jello", 0)
	fmt.Println("hello exists,", c.Exists("hello"))

	v2 := c.Get("key2")
	fmt.Println("key2", v2)

}

func main() {
	fmt.Println("Hello World")
	ExampleClient()
}
*/
