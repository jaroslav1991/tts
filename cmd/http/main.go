package main

import (
	"log"
	"net/http"

	"github.com/jaroslav1991/tts/internal/service"
	serviceHttp "github.com/jaroslav1991/tts/internal/service/http"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// todo
	http.HandleFunc("/", serviceHttp.NewHandler(service.NewService(
		&serviceHttp.HttpDataReader{},
		&service.NoopDataValidator{},
		&service.NoopDataPreparer{},
		&service.NoopDataSaver{},
	)))

	err = http.ListenAndServe(":8484", nil)
}
