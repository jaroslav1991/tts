//go:generate mockgen -source=$GOFILE -destination=interfaces_mocks.go -package=$GOPACKAGE
package dispatcher

import "github.com/jaroslav1991/tts/internal/model"

type Sender interface {
	Send(data []model.DataModel) error
}

type Storage interface {
	ClearSentData(file string) error

	GetFilesToSend() ([]string, error)
	ReadDataToSend(file string) ([]model.DataModel, error)
}
