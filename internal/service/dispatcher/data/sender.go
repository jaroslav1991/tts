package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
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
			_, err = http.Post("/sender", "application/json", bytes.NewBuffer(bytesDataToSend))

			time.Sleep(time.Second)
		}
	}()

	return nil
}
