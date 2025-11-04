package store

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	address := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	db := os.Getenv("REDIS_DB")
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse Redis configuration"))
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       dbInt,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %s", err.Error()))
	}

	fmt.Printf("Redis Pong: %s\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, url string) {
	err := storeService.redisClient.Set(ctx, shortUrl, url, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - url: %s\n", err, shortUrl, url))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	url, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieving key url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return url
}
