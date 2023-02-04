package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jaroslav1991/tts/internal/service/collector"
	"github.com/jaroslav1991/tts/internal/service/model"
)

var (
	ErrInvalidRequestType   = errors.New("expected string")
	ErrUnmarshalRequestData = errors.New("unmarshal request data")
)

type DataReader struct {
	collector.DataReader
}

func (r *DataReader) ReadData(untypedRequest any) (model.DataModel, error) {
	request, ok := untypedRequest.(string)
	if !ok {
		return model.DataModel{}, ErrInvalidRequestType
	}

	var dto DTO

	if err := json.Unmarshal([]byte(request), &dto); err != nil {
		return model.DataModel{}, fmt.Errorf("%w: %v", ErrUnmarshalRequestData, err)
	}

	return model.DataModel{
		Program:  dto.Program,
		Duration: dto.DurationMS * time.Millisecond,
	}, nil
}
