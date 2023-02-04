package data

import "github.com/jaroslav1991/tts/internal/service"

type Validator struct {
	service.DataValidator
}

// ValidateData
// todo implement ValidateData logic
// todo implement tests for ValidateData
func (v Validator) ValidateData(data service.DataModel) error {
	panic("implement me")
}
