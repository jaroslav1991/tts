package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
)

type Sender struct {
	dispatcher.Sender
	HttpAddr string
}

var ErrMarshalData = errors.New("can't marshal data to send")

func (s *Sender) Send(data []model.DataModel) error {
	bytesDataToSend, err := json.Marshal(NewRemoteRequestDTOFromDataModels(data))
	if err != nil {
		log.Printf("%v: %v", ErrMarshalData, err)
		return fmt.Errorf("%w: %v", ErrMarshalData, err)
	}

	resp, err := http.Post(s.HttpAddr, "application/json", bytes.NewBuffer(bytesDataToSend))
	if err != nil {
		log.Println(err)
		return err
	}

	return resp.Body.Close()
}
