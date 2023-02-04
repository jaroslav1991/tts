package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jaroslav1991/tts/internal/service"
)

var (
	ErrInvalidRequestType   = errors.New("expected http.Request")
	ErrReadBodyFailed       = errors.New("read body failed")
	ErrInvalidRequestMethod = errors.New("expected http POST method")
	ErrUnmarshalRequestData = errors.New("unmarshal request data")
)

type DataReader struct {
	service.DataReader
}

func (r *DataReader) ReadData(untypedRequest any) (service.DataModel, error) {
	request, ok := untypedRequest.(*http.Request)
	if !ok {
		return service.DataModel{}, ErrInvalidRequestType
	}

	if request.Method != http.MethodPost {
		return service.DataModel{}, ErrInvalidRequestMethod
	}

	b, err := io.ReadAll(request.Body)
	if err != nil {
		return service.DataModel{}, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	var dto DTO
	if err := json.Unmarshal(b, &dto); err != nil {
		return service.DataModel{}, fmt.Errorf("%w: %v", ErrUnmarshalRequestData, err)
	}

	return service.DataModel{
		Program:  dto.Program,
		Duration: dto.DurationMS * time.Millisecond,
	}, nil
}
