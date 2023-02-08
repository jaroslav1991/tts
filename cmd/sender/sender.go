package main

import (
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"github.com/jaroslav1991/tts/internal/service/dispatcher/data"
	"log"
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
		&data.Storage{},
	)

	err = newService.SendData()
}
