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

		var requestDTO sender.RemoteRequestDTO
		assert.NoError(t, json.Unmarshal(requestBody, &requestDTO))

		if assert.Len(t, requestDTO, 1) {
			requestDTO[0].Events[0].Branch = "some-branch"
			assert.NotEmpty(t, requestDTO[0].Events[0].Branch)

			assert.Equal(t, sender.RemoteRequestDTO{{
				Uid:           "qwerty123",
				PluginType:    "jetbrains",
				PluginVersion: "1.0.0",
				IdeType:       "Intellij idea",
				IdeVersion:    "2.1.1",
				Events: []sender.DTOEvents{
					{
						CreatedAt:      "2022-01-11 14:23:01",
						Type:           "modify file",
						Project:        "some project",
						ProjectBaseDir: "",
						Language:       "golang",
						Target:         "C:/Users/vladimir/IdeaProjects/untitled/.idea/misc.xml",
						Branch:         "some-branch",
					},
				},
			}}, requestDTO)
		}
	}))
	cmd := exec.Command(
		"go",
		"run",
		"../cmd/cli/main.go",
		"-s",
		server.URL,
		"-d",
		`{"uid":"qwerty123","pluginType":"jetbrains","pluginVersion":"1.0.0","ideType":"Intellij idea","ideVersion":"2.1.1","events":[{"createdAt":"2022-01-11 14:23:01","type":"modify file","project":"some project","projectBaseDir":"","language":"golang","target":"C:/Users/vladimir/IdeaProjects/untitled/.idea/misc.xml"}]}`,
	)

	out, err := cmd.CombinedOutput()
	if !assert.NoError(t, err) {
		t.Log("program error output:", string(out))
	}
}
