package dtos

type CreateShortUrlDto struct {
	LongUrl string `json:"long-url" binding:"required"`
}
