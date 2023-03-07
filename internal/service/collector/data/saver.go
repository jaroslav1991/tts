package data

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrCantOpenFile      = errors.New("can't open file")
	ErrCantWriteDataFile = errors.New("can't write data to file")
)

type Saver struct {
	collector.DataSaver
	NewStatsFileName string
}

func (s *Saver) SaveData(data []byte) error {
	file, err := os.OpenFile(s.NewStatsFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Printf("%v: %v", ErrCantOpenFile, err)
		return fmt.Errorf("%w: %v", ErrCantOpenFile, err)
	}

	defer file.Close()

	_, err = file.WriteString(string(data) + "\n")
	if err != nil {
		log.Printf("%v: %v", ErrCantWriteDataFile, err)
		return fmt.Errorf("%w: %v", ErrCantWriteDataFile, err)
	}

	return nil
}
