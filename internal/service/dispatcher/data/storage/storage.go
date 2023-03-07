package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
)

var (
	currentTime      = time.Now
	ErrUnmarshalData = errors.New("can't unmarshal read data")
)

type Storage struct {
	dispatcher.Storage
	NewStatsFileName string
	FilePath         string
}

func (s *Storage) FixDataToSend() (string, error) {
	if err := os.Mkdir(s.FilePath, os.ModePerm); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Printf("can't create path: %v, %v", s.FilePath+string(os.PathSeparator), err)
			return "", fmt.Errorf("can't create path: %v, %w", s.FilePath+string(os.PathSeparator), err)
		}
	}

	nowUnixNano := currentTime().UnixNano()
	newFileName := fmt.Sprintf("%s%d", s.FilePath+string(os.PathSeparator), nowUnixNano)

	if err := os.Rename(s.NewStatsFileName, newFileName); err != nil {
		return "", err
	}
	return newFileName, nil
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
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}

	lines := strings.Split(string(readData), "\n")

	var dataModels []model.DataModel

	for _, line := range lines {
		var dataModel model.DataModel

		if strings.TrimSpace(line) != "" {
			if err := json.Unmarshal([]byte(line), &dataModel); err != nil {
				log.Printf("%v: %v", ErrUnmarshalData, err)
				return nil, fmt.Errorf("%w: %v", ErrUnmarshalData, err)
			}
			dataModels = append(dataModels, dataModel)
		}
	}

	return dataModels, nil
}
