package repository

import (
	"context"
	"fmt"
	. "peterchu999/url-shorterner/database"
	"time"
)

const REDIS_KEY = "shortUrl"
const KEY_TTL = time.Hour * 5

func GetRedisLongUrl(shortUrl string) (string, error) {
	key := fmt.Sprintf("%s:%s", REDIS_KEY, shortUrl)
	res, err := RedisClient.Get(context.TODO(), key).Result()
	return res, err
}

func SetRedisLongUrl(shortUrl string, longUrl string) error {
	key := fmt.Sprintf("%s:%s", REDIS_KEY, shortUrl)
	res := RedisClient.Set(context.TODO(), key, longUrl, KEY_TTL)
	return res.Err()
}
