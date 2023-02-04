package data

import "github.com/jaroslav1991/tts/internal/service"

type Saver struct {
	service.DataSaver
}

// SaveData
// todo implement SaveData logic
// todo implement test for SaveData
func (s *Saver) SaveData([]byte) error {
	panic("implement me")
}
