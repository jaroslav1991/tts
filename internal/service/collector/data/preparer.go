package data

import (
	"encoding/json"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

type Preparer struct {
	collector.DataPreparer
}

func (p Preparer) PrepareData(data model.DataModel) ([]byte, error) {
	return json.Marshal(data)
}
