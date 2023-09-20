package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrUnmarshalData = errors.New("can't unmarshal read data")
)

type Storage struct {
	dispatcher.Storage
	FilePath string
}

func (s *Storage) ClearSentData(file string) error {
	return os.Remove(file)
}

func (s *Storage) GetFilesToSend() ([]string, error) {
	absolutePath, err := filepath.Abs(s.FilePath)
	if err != nil {
		log.Println("can't find absolute path", err)
	}

	files, err := os.ReadDir(s.FilePath)
	if err != nil {
		return nil, err
	}

	filesToSend := []string{}

	for _, file := range files {
		filesToSend = append(filesToSend, absolutePath+string(os.PathSeparator)+file.Name())
	}

	return filesToSend, nil
}

func (s *Storage) ReadDataToSend(file string) ([]model.DataModel, error) {
	readData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(readData), "\n")

	var dataModels []model.DataModel

	for _, line := range lines {
		var dataModel model.DataModel

		if strings.TrimSpace(line) != "" {
			if err := json.Unmarshal([]byte(line), &dataModel); err != nil {
				return nil, fmt.Errorf("%w: %v", ErrUnmarshalData, err)
			}
			dataModels = append(dataModels, dataModel)
		}
	}

	return dataModels, nil
}
