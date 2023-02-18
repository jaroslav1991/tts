package main

import (
	"flag"
	"log"

	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/sender"
	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/storage"
)

var (
	tmpFileName = flag.String(
		"tmpFile",
		"./tempFile",
		"File for temporary storage of stats",
	)
	pathFileName = flag.String(
		"pathToSend",
		"./fileToSend",
		"File for sending to server",
	)
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	newDispatcher := dispatcher.NewService(
		&sender.Sender{HttpAddr: "http://localhost:8080/events"},
		&storage.Storage{
			NewStatsFileName: *tmpFileName,
			FilePath:         *pathFileName,
		},
	)

	err = newDispatcher.SendData()
}
