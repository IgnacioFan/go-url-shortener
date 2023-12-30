package service_test

import (
	"go-url-shortener/internal/service/url_service"
	"go-url-shortener/test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type Expected struct {
  Res interface{}
  Err error
}

func TestGenerateShortURL(t *testing.T) {
  ctrl := gomock.NewController(t)
  zooKeeperMock := mocks.NewMockZookeeper(ctrl)
  urlRepoMock := mocks.NewMockUrlRepository(ctrl)
  service := url_service.Impl{
    ZkClient: zooKeeperMock,
    Repo: urlRepoMock,
    RangeStart: 1,
    RangeEnd: 5,
  }

  tests := []struct {
    Name     string
    Input    string
    RunMock  func()
    Expected *Expected
  }{
    {
      "when success",
      "https://example.com/foobar",
      func() {
        zooKeeperMock.EXPECT().SetNewRange().Return(6,10,nil) 
      },
      &Expected{
        Res: []string{"B","C","D","E","F","G","H"},
        Err: nil,
      },
    },
  }
  for _, test := range tests {
    t.Run(test.Name, func(t *testing.T) {

      test.RunMock()

      for _, v := range test.Expected.Res.([]string) {
        urlRepoMock.EXPECT().Create(test.Input, v).Return(nil)

        res, err := service.GenerateShortURL(test.Input)
        assert.Equal(t, v, res)
        assert.Equal(t, test.Expected.Err, err)
      }
    })
  }
}

func TestShortUrlRedirect(t *testing.T) {
  ctrl := gomock.NewController(t)
  zooKeeperMock := mocks.NewMockZookeeper(ctrl)
  urlRepoMock := mocks.NewMockUrlRepository(ctrl)
  service := url_service.Impl{
    ZkClient: zooKeeperMock,
    Repo: urlRepoMock,
    RangeStart: 1,
    RangeEnd: 5,
  }

  tests := []struct {
    Name     string
    Input    string
    RunMock  func()
    Expected *Expected
  }{
    {
      "when success",
      "B",
      func() {
        urlRepoMock.EXPECT().FindBy("B").Return("https://example.com/foobar", nil)
      },
      &Expected{
        Res: "https://example.com/foobar",
        Err: nil,
      },
    },
  }
  for _, test := range tests {
    t.Run(test.Name, func(t *testing.T) {
      test.RunMock()

      res, err := service.OriginalURL(test.Input)
      assert.Equal(t, test.Expected.Res, res)
      assert.Equal(t, test.Expected.Err, err)
    })
  }
}
