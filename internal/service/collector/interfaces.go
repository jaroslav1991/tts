//go:generate mockgen -source=$GOFILE -destination=interfaces_mocks.go -package=$GOPACKAGE
package collector

import "github.com/jaroslav1991/tts/internal/model"

type DataReader interface {
	ReadData(request any) (model.PluginInfo, error)
}

type DataValidator interface {
	ValidateData(data model.PluginInfo) error
}

type DataAggregator interface {
	Aggregate(info model.PluginInfo) (model.AggregatorInfo, error)
}

type DataPreparer interface {
	PrepareData(
		pluginInfo model.PluginInfo,
		aggregatorInfo model.AggregatorInfo,
	) ([]byte, error)
}

type DataSaver interface {
	SaveData([]byte) error
}
