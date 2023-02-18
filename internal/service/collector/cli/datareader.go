package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrInvalidRequestType   = errors.New("expected string")
	ErrUnmarshalRequestData = errors.New("unmarshal request data")
)

type DataReader struct {
	collector.DataReader
}

func (r *DataReader) ReadData(untypedRequest any) (model.PluginInfo, error) {
	request, ok := untypedRequest.(string)
	if !ok {
		return model.PluginInfo{}, ErrInvalidRequestType
	}

	var dto DTO

	if err := json.Unmarshal([]byte(request), &dto); err != nil {
		return model.PluginInfo{}, fmt.Errorf("%w: %v", ErrUnmarshalRequestData, err)
	}

	return model.PluginInfo{
		Program:     dto.Program,
		Duration:    dto.DurationMS * time.Millisecond,
		PathProject: dto.PathProject,
	}, nil
}
