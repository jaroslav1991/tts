//go:generate mockgen -source=$GOFILE -destination=interfaces_mocks.go -package=$GOPACKAGE
package dispatcher

import "github.com/jaroslav1991/tts/internal/model"

type Sender interface {
	Send(data []model.PluginInfo) error
}

type Storage interface {
	FixDataToSend() (string, error)
	ClearSentData(file string) error

	GetFilesToSend() ([]string, error)
	ReadDataToSend(file string) ([]model.PluginInfo, error)
}
