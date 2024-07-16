package services

import (
	. "peterchu999/url-shorterner/dtos"
	"peterchu999/url-shorterner/model"
	repo "peterchu999/url-shorterner/repository"
	urlUtils "peterchu999/url-shorterner/utils/urls"
	"sync"
)

type URLService struct {
	sync.Mutex
}

var URLServiceObject *URLService

func (us *URLService) CreateShortUrl(createShortUrlDto CreateShortUrlDto) (string, error) {
	us.Lock()
	defer us.Unlock()
	idx, err := repo.GetCurrentCount()
	if err != nil {
		return "", err
	}
	shortUrl := urlUtils.GenerateShortUrl(idx)
	err = repo.CreateUrlData(model.URL{
		LongUrl:  createShortUrlDto.LongUrl,
		ShortUrl: shortUrl,
	})

	return shortUrl, err
}

func init() {
	URLServiceObject = &URLService{}
}
