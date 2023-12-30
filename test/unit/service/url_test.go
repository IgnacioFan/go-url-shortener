package test

import (
	"go-url-shortener/internal/service/url"
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
  zooKeeperMock.EXPECT().SetNewRange().Return(1,5,nil) 
  service, _ := url.InitUrl(zooKeeperMock)

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
        res, err := service.GenerateShortURL(test.Input)
        assert.Equal(t, v, res)
        assert.Equal(t, test.Expected.Err, err)
      }
    })
  }
}
