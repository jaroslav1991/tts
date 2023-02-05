package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jaroslav1991/tts/internal/service/collector"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
	serviceHttp "github.com/jaroslav1991/tts/internal/service/collector/http"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(httpTestsSuite))
}

type httpTestsSuite struct {
	suite.Suite
	tempFile *os.File
	server   *httptest.Server
}

func (s *httpTestsSuite) SetupTest() {
	f, err := os.CreateTemp(os.TempDir(), "testfile")
	if !assert.NoError(s.T(), err) {
		return
	}

	s.tempFile = f
	s.server = httptest.NewServer(serviceHttp.NewHandler(collector.NewService(
		&serviceHttp.DataReader{},
		&data.Validator{},
		&data.Preparer{},
		&data.Saver{
			NewStatsFileName: s.tempFile.Name(),
		},
	)))
}

func (s *httpTestsSuite) TearDownTest() {
	assert.NoError(s.T(), s.tempFile.Close())
	s.server.Close()
}

func (s *httpTestsSuite) TestHttp_Positive() {
	request, err := http.NewRequest(
		http.MethodPost,
		s.server.URL,
		bytes.NewReader([]byte(`
			{
				"program": "IDE1",
				"durationMS": 15000
			}
		`)),
	)
	s.NoError(err)

	response, err := s.server.Client().Do(request)
	s.NoError(err)
	defer response.Body.Close()

	assert.Equal(s.T(), http.StatusOK, response.StatusCode)

	actualData, err := io.ReadAll(s.tempFile)
	s.NoError(err)
	s.NotEqual(`{"Program":"IDE1","Duration":15000000000}`+"\n", string(actualData))
}

func (s *httpTestsSuite) TestHttp_Negative() {
	request, err := http.NewRequest(
		http.MethodPost,
		s.server.URL,
		bytes.NewReader([]byte(``)),
	)
	s.NoError(err)

	response, err := s.server.Client().Do(request)
	s.NoError(err)
	defer response.Body.Close()

	s.Equal(http.StatusInternalServerError, response.StatusCode)

	actualData, err := io.ReadAll(s.tempFile)
	s.NoError(err)
	s.Empty(string(actualData))
}
