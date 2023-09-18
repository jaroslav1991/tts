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
			Id:             event.Id,
			CreatedAt:      event.CreatedAt,
			Type:           event.Type,
			Project:        event.Project,
			ProjectBaseDir: event.ProjectBaseDir,
			Language:       event.Language,
			Target:         event.Target,
			Branch:         event.Branch,
			Timezone:       event.Timezone,
			Params:         event.Params,
		})
	}

	return model.PluginInfo{
		Events: modelEvents,
	}, nil
}
