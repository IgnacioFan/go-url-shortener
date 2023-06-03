package integration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HealthTestSuite struct {
	TestSuite
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}

func (t *HealthTestSuite) TestHealth() {
	res, err := http.Get(fmt.Sprintf("%s/health", t.TestServer.URL))
	t.Require().NoError(err)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	t.Require().Equal(http.StatusOK, res.StatusCode)
	t.Require().Equal(`{"message":"healthy","data":null}`, string(body))
}
