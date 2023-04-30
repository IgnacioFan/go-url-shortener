package shorturl

import (
	"errors"
	"go-url-shortener/internal/entity"
	"go-url-shortener/internal/usecase/base62"
	"go-url-shortener/pkg/redis"
)

type ShortUrl struct {
	Client redis.RedisClient
	Repo   entity.ShortUrlRepository
}

func NewShortUrlUsecase(repo entity.ShortUrlRepository, client redis.RedisClient) entity.ShortUrlUsecase {
	return &ShortUrl{
		Client: client,
		Repo:   repo,
	}
}

func (i *ShortUrl) Create(url string) (string, error) {
	id, err := i.Repo.Create(url)
	if err != nil {
		return "", err
	}
	return base62.Encode(id), nil
}

func (i *ShortUrl) Redirect(code string) (string, error) {
	id, err := base62.Decode(code)
	if err != nil {
		return "", err
	}

	url, err := i.Client.Get(code)
	if len(url) != 0 && err == nil {
		return url, err
	} else if len(url) == 0 && err == nil {
		return "", errors.New("URL not found.")
	} else if err != nil && err.Error() == "No entry" {
		url, err := i.Repo.Find(id)
		_ = i.Client.Set(code, url)
		return url, err
	} else {
		return "", err
	}
}

func (i *ShortUrl) Delete(code string) error {
	id, err := base62.Decode(code)
	if err != nil {
		return err
	}
	res, err := i.Repo.Delete(id)
	if res == 1 && err == nil {
		// print out if failure
		_ = i.Client.Del(code)
		return nil
	} else if res == 0 && err == nil {
		return errors.New("URL not found.")
	} else {
		return err
	}
}
