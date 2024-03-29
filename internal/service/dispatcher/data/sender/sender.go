package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
)

type Sender struct {
	dispatcher.Sender
	HttpAddr string
	AuthKey  string
}

var ErrMarshalData = errors.New("can't marshal data to send")

func (s *Sender) Send(data []model.DataModel) error {
	bytesDataToSend, err := json.Marshal(NewRemoteRequestDTOFromDataModels(data))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrMarshalData, err)
	}

	req, err := http.NewRequest("POST", s.HttpAddr, bytes.NewBuffer(bytesDataToSend))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", s.AuthKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}
