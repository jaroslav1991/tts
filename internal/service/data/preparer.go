package data

import (
	"encoding/json"
	"github.com/jaroslav1991/tts/internal/service"
)

type Preparer struct {
	service.DataPreparer
}

func (p Preparer) PrepareData(data service.DataModel) ([]byte, error) {
	return json.Marshal(data)
}
