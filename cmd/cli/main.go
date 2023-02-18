package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	"github.com/jaroslav1991/tts/internal/service/collector/data/aggregator"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	dataDis "github.com/jaroslav1991/tts/internal/service/dispatcher/data"

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
		"d",
		"",
		"Stats data JSON string",
	)

	pathFileName = flag.String(
		"pathToSend",
		"./fileToSend",
		"File for sending to server",
	)

	httpRemote = flag.String(
		"s",
		"http://localhost:8080/events",
		"Http address for sending events",
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

	newCollector := collector.NewService(
		&cli.DataReader{},
		&data.Validator{},
		&data.Aggregator{
			Aggregators: []data.MergeAggregator{
				&aggregator.CurrentBranchAggregator{},
			},
		},
		&data.Preparer{},
		&data.Saver{NewStatsFileName: *tmpFileName},
	)

	newDispatcher := dispatcher.NewService(
		&dataDis.Sender{HttpAddr: *httpRemote},
		&dataDis.Storage{
			NewStatsFileName: *tmpFileName,
			FilePath:         *pathFileName,
		},
	)

	// todo собирать информацию о проекте
	if err = newCollector.SaveData(*inputData); err != nil {
		return
	}

	err = newDispatcher.SendData()
}
