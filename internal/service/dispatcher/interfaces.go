package dispatcher

import "github.com/jaroslav1991/tts/internal/service/model"

type Sender interface {
	Send(data []model.DataModel) error
}

type Storage interface {
	FixDataToSend() error
	ClearSentData(file string) error

	GetFilesToSend() ([]string, error)
	ReadDataToSend(file string) ([]model.DataModel, error)
}
