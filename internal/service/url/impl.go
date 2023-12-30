package url

import (
	"go-url-shortener/internal/adpater/zookeeper"
	"go-url-shortener/internal/service/base62"
	"sync"
)

var tokenLock sync.Mutex

type UrlService interface {
  GenerateShortURL(longUrl string) (string, error)
  OriginalURL(shortUrl string) (string, error)
}

type Impl struct {
  ZkClient zookeeper.Zookeeper
  RangeStart int
  RangeEnd int
}

func InitUrl(zkClient zookeeper.Zookeeper) (*Impl, error) {
  start, end, err := zkClient.SetNewRange()
  if err != nil {
    return nil, err
  }
  return &Impl{
    ZkClient: zkClient,
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
  // TODO: store it into the database
  return shortUrl, nil
}

func (i *Impl) OriginalURL(shortURL string) (string, error)  {
  return "https://example.com/foobar", nil
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
