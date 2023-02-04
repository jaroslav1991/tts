package service

import (
	"fmt"
)

func NewService(reader DataReader, validator DataValidator, preparer DataPreparer, saver DataSaver) *Service {
	return &Service{reader: reader, validator: validator, preparer: preparer, saver: saver}
}

type Service struct {
	reader    DataReader
	validator DataValidator
	preparer  DataPreparer
	saver     DataSaver
}

func (s *Service) SaveData(request any) error {
	data, err := s.reader.ReadData(request)
	if err != nil {
		return fmt.Errorf("read data from request failed: %w", err)
	}

	if err := s.validator.ValidateData(data); err != nil {
		return fmt.Errorf("validate data failed: %w", err)
	}

	dataForSave, err := s.preparer.PrepareData(data)
	if err != nil {
		return fmt.Errorf("prepare data for save failed: %w", err)
	}

	if err := s.saver.SaveData(dataForSave); err != nil {
		return fmt.Errorf("save data failed: %w", err)
	}

	return nil
}

// SendData
// todo implement SendData logic
// todo implement tests for SendData
func (s *Service) SendData() error {
	panic("implement me")
}
