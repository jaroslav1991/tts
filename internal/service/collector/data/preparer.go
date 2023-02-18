package data

import (
	"encoding/json"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

type Preparer struct {
	collector.DataPreparer
}

func (p Preparer) PrepareData(
	pluginInfo model.PluginInfo,
	aggregationInfo model.AggregatorInfo,
) ([]byte, error) {
	return json.Marshal(model.DataModel{
		PluginInfo:     pluginInfo,
		AggregatorInfo: aggregationInfo,
	})
}
