package data

import (
	"fmt"
	"os"
	"time"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
)

var currentTime = time.Now

type Storage struct {
	dispatcher.Storage
	NewStatsFileName string
}

func (s *Storage) FixDataToSend() (string, error) {
	nowUnixNano := currentTime().UnixNano()
	newFileName := fmt.Sprintf("%s.%d", s.NewStatsFileName, nowUnixNano)

	if err := os.Rename(s.NewStatsFileName, newFileName); err != nil {
		return "", err
	}
	return newFileName, nil
}

func (s *Storage) ClearSentData(file string) error {
	return os.Remove(file)
}

func (s *Storage) GetFilesToSend() ([]string, error) {
	//TODO implement GetFilesToSend
	panic("implement me")
}

func (s *Storage) ReadDataToSend(file string) ([]model.DataModel, error) {
	//TODO implement ReadDataToSend
	panic("implement me")
}
