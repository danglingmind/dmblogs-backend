// +build integration

package integration_tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"

	"github.com/arielizuardi/golang-backend-blog/config"
)

type e2eTestSuite struct {
	suite.Suite
	port int
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{port: 8000})
}

func (s *e2eTestSuite) SetupSuite() {
	s.Require().NoError(config.Load())
}

func (s *e2eTestSuite) TearDownSuite() {
}

func (s *e2eTestSuite) SetupTest() {
}

func (s *e2eTestSuite) TearDownTest() {
}

func (s *e2eTestSuite) Test_Integration_VerifyThatLoginActionReturnsNoRecordFoundForNonExistingUser() {

	body := strings.NewReader(`{
		"email" : "deepak@gmail.com",
		"password" : "test"
	}`)
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/users/login", s.port), body)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusForbidden, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"reason":"record not found"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()

}
