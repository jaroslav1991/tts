package main

import (
	"flag"
	"github.com/jaroslav1991/tts/internal/service/collector"
	"github.com/jaroslav1991/tts/internal/service/collector/cli"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
	"github.com/jaroslav1991/tts/internal/service/collector/data/aggregator"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/sender"
	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/storage"
	"log"
	"os"
	"strings"
)

var (
	tmpFileName = flag.String(
		"t",
		"./stats",
		"File for temporary storage of stats",
	)

	inputData = flag.String(
		"d",
		"",
		"Stats data in JSON format string",
	)

	pathToSendingFiles = flag.String(
		"o",
		"./outbox",
		"Path for temporary store files",
	)

	httpRemote = flag.String(
		"s",
		"http://localhost:8080/events",
		"Http address for sending events",
	)

	authKey = flag.String(
		"k",
		"",
		"authorization key",
	)
)

func init() {
	fileInfo, err := os.OpenFile("cli-logger.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	log.SetOutput(fileInfo)

}

func main() {
	flag.Parse()

	log.Println("CLI starting...")

	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if strings.TrimSpace(*inputData) == "" {
		flag.Usage()
		return
	}

	// newCollector
	// - receive data from plugin
	// - aggregate advanced data
	newCollector := collector.NewService(
		&cli.DataReader{},
		&data.Validator{},
		&data.Aggregator{
			Aggregators: []data.MergeAggregator{
				&aggregator.CommonAggregator{},
			},
		},
		&data.Preparer{},
		&data.Saver{NewStatsFileName: *tmpFileName},
	)

	newDispatcher := dispatcher.NewService(
		&sender.Sender{HttpAddr: *httpRemote, AuthKey: *authKey},
		&storage.Storage{
			NewStatsFileName: *tmpFileName,
			FilePath:         *pathToSendingFiles,
		},
	)

	if err = newCollector.SaveData(*inputData); err != nil {
		return
	}

	err = newDispatcher.SendData()
	if err == nil {
		log.Println("sending success")
	}
}
