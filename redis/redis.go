package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var Instance *redis.Client

var once sync.Once

func Connect() {
	once.Do(func() {
		Instance = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			PoolSize: 20,
		})
	})
}

func Close() {
	Instance.Close()
}
