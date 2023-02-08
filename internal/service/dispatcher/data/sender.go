package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"log"
	"net/http"
	"time"
)

type Sender struct {
	dispatcher.Sender
}

var ErrMarshalData = errors.New("can't marshal data to send")

func (s *Sender) Send(data []model.DataModel) error {
	bytesDataToSend, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshalData, err)
	}

	go func() {
		for {
			resp, err := http.Post("/sender", "application/json", bytes.NewBuffer(bytesDataToSend))
			if err != nil {
				return
			}

			err = resp.Body.Close()
			if err != nil {
				log.Println("can't close response body", err)
				return
			}

			time.Sleep(time.Second * 60)
		}
	}()

	return nil
}
