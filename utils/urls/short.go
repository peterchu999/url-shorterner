package urls

import "math"

func GenerateShortUrl(n int) string {
	dict := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var shortUrl []byte
	for !(n < 1) {
		shortUrl = append(shortUrl, dict[n%62])
		n = int(math.Floor(float64(n / 62)))
	}

	return string(shortUrl)
}
