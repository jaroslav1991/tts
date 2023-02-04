package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	"github.com/jaroslav1991/tts/internal/service"
	"github.com/jaroslav1991/tts/internal/service/cli"
	"github.com/jaroslav1991/tts/internal/service/data"
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

	newService := service.NewService(
		&cli.DataReader{},
		&data.Validator{},
		&data.Preparer{},
		&data.Saver{FileName: *tmpFileName},
	)

	err = newService.SaveData(*inputData)
}
