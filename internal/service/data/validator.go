package data

import (
	"errors"
	"github.com/jaroslav1991/tts/internal/service"
	"strings"
)

var (
	ErrInvalidProgramField  = errors.New("invalid program field")
	ErrInvalidDurationField = errors.New("invalid duration field")
)

type Validator struct {
	service.DataValidator
}

func (v Validator) ValidateData(data service.DataModel) error {

	if strings.TrimSpace(data.Program) == "" {
		return ErrInvalidProgramField
	}

	if data.Duration < 1 {
		return ErrInvalidDurationField
	}

	return nil
}
