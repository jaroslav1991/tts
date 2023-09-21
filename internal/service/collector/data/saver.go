package data

import (
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/service/collector"
	"log"
	"os"
)

var (
	ErrCantOpenFile      = errors.New("can't open file")
	ErrCantCreatePath    = errors.New("can't create path")
	ErrCantWriteDataFile = errors.New("can't write data to file")
)

type Saver struct {
	collector.DataSaver
	NewStatsFileName string
	AuthKey          string
}

func (s *Saver) SaveData(data []byte) error {
	if err := os.Mkdir(s.NewStatsFileName, os.ModePerm); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Printf("%v, %v", ErrCantCreatePath, err)
			return fmt.Errorf("%w, %v", ErrCantCreatePath, err)
		}
	}

	newFileName := fmt.Sprintf("%s%s", s.NewStatsFileName+string(os.PathSeparator), s.AuthKey)

	file, err := os.OpenFile(newFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCantOpenFile, err)

	}

	defer file.Close()

	_, err = file.WriteString(string(data) + "\n")
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCantWriteDataFile, err)
	}

	return nil
}
