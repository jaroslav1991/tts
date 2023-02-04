package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jaroslav1991/tts/internal/service"
)

type HttpDataReader struct {
}

func (r *HttpDataReader) ReadData(untypedRequest any) (service.DataModel, error) {
	request, ok := untypedRequest.(*http.Request)
	if !ok {
		return service.DataModel{}, errors.New("expected http.Request")
	}

	b, err := io.ReadAll(request.Body)
	if err != nil {
		return service.DataModel{}, fmt.Errorf("read body: %w", err)
	}

	var result service.DataModel

	if err := json.Unmarshal(b, &result); err != nil {
		return service.DataModel{}, fmt.Errorf("marshal request data: %w", err)
	}

	return result, nil
}
