package collector

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func NewService(
	reader DataReader,
	validator DataValidator,
	aggregator DataAggregator,
	preparer DataPreparer,
	saver DataSaver,
) *Service {
	return &Service{
		reader:     reader,
		validator:  validator,
		aggregator: aggregator,
		preparer:   preparer,
		saver:      saver,
		uudGenFn:   uuid.NewString,
	}
}

type Service struct {
	reader     DataReader
	validator  DataValidator
	aggregator DataAggregator
	preparer   DataPreparer
	saver      DataSaver
	uudGenFn   func() string
}

func (s *Service) SaveData(request any) error {
	pluginInfo, err := s.reader.ReadData(request)
	if err != nil {
		return fmt.Errorf("read pluginInfo from request failed: %w", err)
	}

	if err := s.validator.ValidateData(pluginInfo); err != nil {
		return fmt.Errorf("validate pluginInfo failed: %w", err)
	}

	for i := range pluginInfo.Events {
		pluginInfo.Events[i].Uid = s.uudGenFn()
	}

	aggregated, err := s.aggregator.Aggregate(pluginInfo)
	if err != nil {
		return fmt.Errorf("aggregation failed: %w", err)
	}

	dataForSave, err := s.preparer.PrepareData(pluginInfo, aggregated)
	if err != nil {
		log.Printf("prepare pluginInfo for save failed: %v", err)
		return fmt.Errorf("prepare pluginInfo for save failed: %w", err)
	}

	if err := s.saver.SaveData(dataForSave); err != nil {
		return fmt.Errorf("save pluginInfo failed: %w", err)
	}

	return nil
}
