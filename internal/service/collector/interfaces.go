//go:generate mockgen -source=$GOFILE -destination=interfaces_mocks.go -package=$GOPACKAGE
package collector

import "github.com/jaroslav1991/tts/internal/service/model"

type DataReader interface {
	ReadData(request any) (model.DataModel, error)
}

type DataValidator interface {
	ValidateData(data model.DataModel) error
}

type DataPreparer interface {
	PrepareData(data model.DataModel) ([]byte, error)
}

type DataSaver interface {
	SaveData([]byte) error
}
