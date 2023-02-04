package cli

import (
	"github.com/jaroslav1991/tts/internal/service"
)

type HttpDataReader struct {
}

// ReadData
// todo implement ReadData logic
// todo implement test for ReadData
func (r *HttpDataReader) ReadData(untypedRequest any) (service.DataModel, error) {
	panic("implement me")
}
