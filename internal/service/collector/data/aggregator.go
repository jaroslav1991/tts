package data

import (
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

type MergeAggregator interface {
	Aggregate(info model.PluginInfo, target *model.AggregatorInfo) error
}

type Aggregator struct {
	collector.DataAggregator
	Aggregators []MergeAggregator
}

func (a *Aggregator) Aggregate(info model.PluginInfo) (model.AggregatorInfo, error) {
	var result model.AggregatorInfo
	for _, aggregator := range a.Aggregators {
		if err := aggregator.Aggregate(info, &result); err != nil {
			return model.AggregatorInfo{}, fmt.Errorf("aggregation failed: %w", err)
		}
	}
	return result, nil
}
