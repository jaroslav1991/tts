package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jaroslav1991/tts/internal/service/collector"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
	serviceHttp "github.com/jaroslav1991/tts/internal/service/collector/http"
)

var (
	tmpFileName = flag.String(
		"tmpFile",
		"./tempFile",
		"File for temporary storage of stats",
	)

	httpAddr = flag.String(
		"httpAddr",
		":8484",
		"Http address for handling stats",
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

	http.HandleFunc("/", serviceHttp.NewHandler(collector.NewService(
		&serviceHttp.DataReader{},
		&data.Validator{},
		&data.Preparer{},
		&data.Saver{
			NewStatsFileName: *tmpFileName,
		},
	)))

	err = http.ListenAndServe(*httpAddr, nil)
}
