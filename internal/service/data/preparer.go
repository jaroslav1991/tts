package data

import "github.com/jaroslav1991/tts/internal/service"

type Preparer struct {
	service.DataPreparer
}

// PrepareData
// todo implement PrepareData logic
// todo implement tests for PrepareData
func (p Preparer) PrepareData(data service.DataModel) ([]byte, error) {
	panic("implement me")
}
