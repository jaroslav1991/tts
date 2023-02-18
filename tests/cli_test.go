package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestBody, err := io.ReadAll(request.Body)
		assert.NoError(t, err)

		expected := `[{"program":"some program","duration":15000000,"pathProject":"../","currentGitBranch":"aggregation"}]`
		assert.Equal(t, expected, string(requestBody))
	}))

	cmd := exec.Command(
		"go",
		"run",
		"../cmd/cli/main.go",
		"-s",
		server.URL,
		"-d",
		`{"program": "some program", "durationMS": 15, "pathProject": "../"}`,
	)

	out, err := cmd.CombinedOutput()
	if !assert.NoError(t, err) {
		t.Log("program error output:", string(out))
	}
}
