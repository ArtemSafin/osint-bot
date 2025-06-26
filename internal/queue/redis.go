package queue

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

const QueueName = "email_tasks"

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	log.Println("🔌 Redis подключен")
}

func Push(email string) error {
	return Rdb.RPush(Ctx, QueueName, email).Err()
}

func Pop() (string, error) {
	res, err := Rdb.BLPop(Ctx, 5*time.Second, QueueName).Result()
	if err != nil || len(res) < 2 {
		return "", err
	}
	return res[1], nil
}
