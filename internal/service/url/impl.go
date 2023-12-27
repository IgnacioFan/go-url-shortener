package url

import "go-url-shortener/internal/service/base62"

func GenerateShortURL(longURL string) (string, error) {
	shortUrl := base62.Encode(uint64(uniqueID()))
	// TODO: update the counter
	// TODO: store it into the database
	return shortUrl, nil
}

func OriginalURL(shortURL string) (string, error)  {
	return "https://example.com/foobar", nil
}

// TODO: retreiv value from Zookeeper
func uniqueID() int {
	return 1 
}
