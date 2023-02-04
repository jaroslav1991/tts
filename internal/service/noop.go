package service

type NoopDataReader struct {
}

func (NoopDataReader) ReadData(request any) (DataModel, error) {
	return DataModel{}, nil
}

type NoopDataValidator struct {
}

func (NoopDataValidator) ValidateData(data DataModel) error {
	return nil
}

type NoopDataPreparer struct {
}

func (NoopDataPreparer) PrepareData(data DataModel) ([]byte, error) {
	return nil, nil
}

type NoopDataSaver struct {
}

func (NoopDataSaver) SaveData([]byte) error {
	return nil
}
