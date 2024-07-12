package services

import "peterchu999/url-shorterner/repository"

func GetRedirectUrl(shortUrl string) (string, error) {
	url, err := repository.GetUrlData(shortUrl)
	return url.LongUrl, err
}
