package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrInvalidRequestType   = errors.New("expected http.Request")
	ErrReadBodyFailed       = errors.New("read body failed")
	ErrInvalidRequestMethod = errors.New("expected http POST method")
	ErrUnmarshalRequestData = errors.New("unmarshal request data")
)

type DataReader struct {
	collector.DataReader
}

func (r *DataReader) ReadData(untypedRequest any) (model.PluginInfo, error) {
	request, ok := untypedRequest.(*http.Request)
	if !ok {
		return model.PluginInfo{}, ErrInvalidRequestType
	}

	if request.Method != http.MethodPost {
		return model.PluginInfo{}, ErrInvalidRequestMethod
	}

	b, err := io.ReadAll(request.Body)
	if err != nil {
		return model.PluginInfo{}, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	var dto DTO
	if err := json.Unmarshal(b, &dto); err != nil {
		return model.PluginInfo{}, fmt.Errorf("%w: %v", ErrUnmarshalRequestData, err)
	}

	return model.PluginInfo{
		Program:  dto.Program,
		Duration: dto.DurationMS * time.Millisecond,
	}, nil
}
