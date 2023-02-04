package data

import (
	"errors"
	"strings"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrInvalidProgramField  = errors.New("invalid program field")
	ErrInvalidDurationField = errors.New("invalid duration field")
)

type Validator struct {
	collector.DataValidator
}

func (v Validator) ValidateData(data model.DataModel) error {

	if strings.TrimSpace(data.Program) == "" {
		return ErrInvalidProgramField
	}

	if data.Duration < 1 {
		return ErrInvalidDurationField
	}

	return nil
}
