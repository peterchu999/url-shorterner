package services

import (
	"peterchu999/url-shorterner/repository"
)

func GetRedirectUrl(shortUrl string) (string, error) {
	url, err := repository.GetUrlData(shortUrl)
	return url.LongUrl, err
}

func GetRedirectUrlFast(shortUrl string) (string, error) {
	redisValue, redisErr := repository.GetRedisLongUrl(shortUrl)
	// log.Println("got from redis", redisValue, redisErr)
	if redisErr == nil {
		return redisValue, nil
	}
	// log.Println("not taking from REDISSSSSS")
	url, err := repository.GetUrlData(shortUrl)
	go repository.SetRedisLongUrl(shortUrl, url.LongUrl)
	return url.LongUrl, err
}
