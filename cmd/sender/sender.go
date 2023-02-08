package main

import (
	"flag"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"github.com/jaroslav1991/tts/internal/service/dispatcher/data"
	"log"
)

var (
	tmpFileName = flag.String(
		"tmpFile",
		"./tempFile",
		"File for temporary storage of stats",
	)
	pathFileName = flag.String(
		"fileToSend",
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

	newService := dispatcher.NewService(
		&data.Sender{},
		&data.Storage{
			NewStatsFileName: *tmpFileName,
			FilePath:         *pathFileName,
		},
	)

	err = newService.SendData()
}
