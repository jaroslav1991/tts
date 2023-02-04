package main

import (
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/service"
	"github.com/jaroslav1991/tts/internal/service/cli"
	"github.com/jaroslav1991/tts/internal/service/data"
	"log"
	"os"
)

// todo implement cli logic

//var (
//	tmpFileName = flag.String(
//		"tmpFile",
//		"./tempFile",
//		"File for temporary storage of stats",
//	)
//
//	inputData = flag.String(
//		"data",
//		"",
//		"Stats data",
//	)
//)

func main() {

	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if len(os.Args) < 2 {
		err = errors.New("need more than 1 args")
		return
	}

	newService := service.NewService(&cli.DataReader{}, &data.Validator{}, &data.Preparer{}, &data.Saver{FileName: "./tmpFile"})

	err = newService.SaveData(os.Args[1])

	fmt.Println(os.Args[1])

}
