package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"net/http"
)

type Sender struct {
	dispatcher.Sender
	HttpAddr string
}

var ErrMarshalData = errors.New("can't marshal data to send")

func (s *Sender) Send(data []model.DataModel) error {
	bytesDataToSend, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshalData, err)
	}

	resp, err := http.Post(s.HttpAddr, "application/json", bytes.NewBuffer(bytesDataToSend))
	if err != nil {
		return err
	}

	return resp.Body.Close()
}
