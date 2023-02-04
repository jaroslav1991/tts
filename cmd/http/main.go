package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jaroslav1991/tts/internal/service"
	"github.com/jaroslav1991/tts/internal/service/data"
	serviceHttp "github.com/jaroslav1991/tts/internal/service/http"
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

	http.HandleFunc("/", serviceHttp.NewHandler(service.NewService(
		&serviceHttp.DataReader{},
		&data.Validator{},
		&data.Preparer{},
		&data.Saver{
			FileName: *tmpFileName,
		},
	)))

	err = http.ListenAndServe(*httpAddr, nil)
}
