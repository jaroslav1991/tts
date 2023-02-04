package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/service"
	"time"
)

var (
	ErrInvalidRequestType   = errors.New("expected string")
	ErrUnmarshalRequestData = errors.New("unmarshal request data")
)

type DataReader struct {
	service.DataReader
}

func (r *DataReader) ReadData(untypedRequest any) (service.DataModel, error) {
	request, ok := untypedRequest.(string)
	if !ok {
		return service.DataModel{}, ErrInvalidRequestType
	}

	var dto DTO

	if err := json.Unmarshal([]byte(request), &dto); err != nil {
		return service.DataModel{}, fmt.Errorf("%w: %v", ErrUnmarshalRequestData, err)
	}

	return service.DataModel{
		Program:  dto.Program,
		Duration: dto.DurationMS * time.Millisecond,
	}, nil
}
