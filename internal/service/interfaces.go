//go:generate mockgen -source=$GOFILE -destination=interfaces_mocks.go -package=$GOPACKAGE
package service

type DataReader interface {
	ReadData(request any) (DataModel, error)
}

type DataValidator interface {
	ValidateData(data DataModel) error
}

type DataPreparer interface {
	PrepareData(data DataModel) ([]byte, error)
}

type DataSaver interface {
	SaveData([]byte) error
}
