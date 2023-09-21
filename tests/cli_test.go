package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/sender"
)

func TestCliSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestBody, err := io.ReadAll(request.Body)
		assert.NoError(t, err)

		var requestDTO sender.RemoteRequestDTOItem
		assert.NoError(t, json.Unmarshal(requestBody, &requestDTO))

		assert.Equal(t, "secretKeyToken", request.Header.Get("Authorization"))
	}))
	cmd := exec.Command(
		"go",
		"run",
		"../cmd/cli/main.go",
		"-s",
		server.URL,
		"-d",
		`{"events":[{"id":"qwerty","createdAt":"2022-01-11 14:23:01","type":"modify file","project":"some project","projectBaseDir":"C:/Users/jaros/GolandProjects/tts","language":"golang","target":"C:/Users/vladimir/IdeaProjects/untitled/.idea/misc.xml"}]}`,
		"-k",
		"secretKeyToken",
	)

	out, err := cmd.CombinedOutput()
	if !assert.NoError(t, err) {
		t.Log("program error output:", string(out))
	}
}
