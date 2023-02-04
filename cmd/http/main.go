package main

import (
	"log"
	"net/http"

	"github.com/jaroslav1991/tts/internal/service"
	"github.com/jaroslav1991/tts/internal/service/data"
	serviceHttp "github.com/jaroslav1991/tts/internal/service/http"
)

func main() {
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
		&data.Saver{},
	)))

	err = http.ListenAndServe(":8484", nil)
}
