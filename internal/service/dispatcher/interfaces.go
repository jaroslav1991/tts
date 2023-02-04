package dispatcher

import "github.com/jaroslav1991/tts/internal/model"

// TODO implement http sender to remote server
type Sender interface {
	Send(data []model.DataModel) error
}

// TODO implement storage
type Storage interface {
	FixDataToSend() error
	ClearSentData(file string) error

	GetFilesToSend() ([]string, error)
	ReadDataToSend(file string) ([]model.DataModel, error)
}
