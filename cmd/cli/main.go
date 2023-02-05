package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	"github.com/jaroslav1991/tts/internal/service/collector"
	"github.com/jaroslav1991/tts/internal/service/collector/cli"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
)

var (
	tmpFileName = flag.String(
		"tmpFile",
		"./tempFile",
		"File for temporary storage of stats",
	)

	inputData = flag.String(
		"data",
		"",
		"Stats data JSON string",
	)
)

func main() {
	flag.Parse()

	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if strings.TrimSpace(*inputData) == "" {
		err = errors.New("provide stats data JSON with key -data")
		return
	}

	newService := collector.NewService(
		&cli.DataReader{},
		&data.Validator{},
		&data.Preparer{},
		&data.Saver{NewStatsFileName: *tmpFileName},
	)

	err = newService.SaveData(*inputData)
}
