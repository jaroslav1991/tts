package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
	"log"
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
		log.Println(ErrInvalidRequestType)
		return model.PluginInfo{}, ErrInvalidRequestType
	}

	var dto DTO

	if err := json.Unmarshal([]byte(request), &dto); err != nil {
		log.Printf("%v: %v", ErrUnmarshalRequestData, err)
		return model.PluginInfo{}, fmt.Errorf("%w: %v", ErrUnmarshalRequestData, err)
	}

	var modelEvents []model.Events

	for _, event := range dto.Events {
		modelEvents = append(modelEvents, model.Events{
			Uid:       event.Uid,
			CreatedAt: event.CreatedAt,
			Type:      event.Type,
			Project:   event.Project,
			Language:  event.Language,
			Target:    event.Target,
			Branch:    event.Branch,
			Params:    event.Params,
		})
	}

	return model.PluginInfo{
		PluginType:    dto.PluginType,
		PluginVersion: dto.PluginVersion,
		CliType:       dto.CliType,
		CliVersion:    dto.CliVersion,
		DeviceName:    dto.DeviceName,
		Events:        modelEvents,
	}, nil
}
