package url_service

import (
	"go-url-shortener/internal/adpater/zookeeper"
	"go-url-shortener/internal/repository/url_repo"
	"go-url-shortener/internal/service/base62"
	"go-url-shortener/pkg/postgres"
	"sync"
)

var tokenLock sync.Mutex

type UrlService interface {
  GenerateShortURL(longUrl string) (string, error)
  OriginalURL(shortUrl string) (string, error)
}

type Impl struct {
  ZkClient zookeeper.Zookeeper
  Repo url_repo.UrlRepository
  RangeStart int
  RangeEnd int
}

func NewUrlService() (UrlService, error) {
  zkClient, err := zookeeper.InitZooKeeper()
  if err != nil {
    return nil, err
  }
  db, err := postgres.NewPostgres()
  if err != nil {
    return nil, err
  }
  repo := url_repo.NewShortUrlRepo(db)

  start, end, err := zkClient.SetNewRange()
  if err != nil {
    return nil, err
  }
  return &Impl{
    ZkClient: zkClient,
    Repo: repo,
    RangeStart: start,
    RangeEnd: end,
  }, nil
}

func (i *Impl) GenerateShortURL(longURL string) (string, error) {
  uniqueId, err := i.GetUniqeId()
  if err != nil {
    return "", err
  }
  shortUrl := base62.Encode(uint64(uniqueId))
	if err := i.Repo.Create(longURL, shortUrl); err != nil {
		return "", err
	}
  return shortUrl, nil
}

func (i *Impl) OriginalURL(shortURL string) (string, error)  {
  longURL, err := i.Repo.FindBy(shortURL)
  if ; err != nil {
		return "", err
	}
  return longURL, nil
}

func (i *Impl) GetUniqeId() (int, error) {
  var res int
  tokenLock.Lock()
  res = i.RangeStart
  if i.RangeStart == i.RangeEnd {
		start, end, err := i.ZkClient.SetNewRange()
		if err != nil {
      return -1, err
    }
    i.RangeStart = start
    i.RangeEnd = end
	} else {
    i.RangeStart++
  }
  tokenLock.Unlock()
  return res, nil
}
