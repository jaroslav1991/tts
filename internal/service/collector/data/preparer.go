package data

import (
	"encoding/json"

	"github.com/jaroslav1991/tts/internal/service/collector"
	"github.com/jaroslav1991/tts/internal/service/model"
)

type Preparer struct {
	collector.DataPreparer
}

func (p Preparer) PrepareData(data model.DataModel) ([]byte, error) {
	return json.Marshal(data)
}
