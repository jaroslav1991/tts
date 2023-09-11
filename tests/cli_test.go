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

		assert.Equal(t, "secretKeyToken", request.Header.Get("Authorization"))

		if assert.Len(t, requestDTO, 1) {
			requestDTO[0].Events[0].Branch = "some-branch"
			requestDTO[0].OSName = "windows"
			requestDTO[0].Events[0].Id = "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84"
			assert.NotEmpty(t, requestDTO[0].Events[0].Branch)
			assert.NotEmpty(t, requestDTO[0].OSName)
			assert.NotEmpty(t, requestDTO[0].Events[0].Id)

			assert.Equal(t, sender.RemoteRequestDTO{{
				PluginType:    "jetbrains",
				PluginVersion: "1.0.0",
				CliType:       "windowsOS",
				CliVersion:    "1.1.0",
				OSName:        "windows",
				IdeType:       "Intellij idea",
				IdeVersion:    "2.1.1",
				Events: []sender.DTOEvents{
					{
						Id:             "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
						CreatedAt:      "2022-01-11 14:23:01",
						Type:           "modify file",
						Project:        "some project",
						ProjectBaseDir: "",
						Language:       "golang",
						Target:         "C:/Users/vladimir/IdeaProjects/untitled/.idea/misc.xml",
						Branch:         "some-branch",
						Timezone:       "123456789",
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
		`{"pluginType":"jetbrains","pluginVersion":"1.0.0","cliType":"windowsOS","cliVersion":"1.1.0","osName":"","ideType":"Intellij idea","ideVersion":"2.1.1","events":[{"id":"","createdAt":"2022-01-11 14:23:01","type":"modify file","project":"some project","projectBaseDir":"","language":"golang","target":"C:/Users/vladimir/IdeaProjects/untitled/.idea/misc.xml","timezone":"123456789"}]}`,
		"-k",
		"secretKeyToken",
	)

	out, err := cmd.CombinedOutput()
	if !assert.NoError(t, err) {
		t.Log("program error output:", string(out))
	}
}
