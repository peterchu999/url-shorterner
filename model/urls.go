package model

type URL struct {
	ShortUrl string `bson:"short_url"`
	LongUrl  string `bson:"long_url"`
}

type URLIndex struct {
	Count int    `bson:"count"`
	Name  string `bson:"name"`
}
