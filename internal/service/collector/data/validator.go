package data

import (
	"errors"
	"strings"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrInvalidCreatedAtFiled = errors.New("invalid created at field")
	ErrInvalidTypeFiled      = errors.New("invalid type field")
)

type Validator struct {
	collector.DataValidator
}

func (v Validator) ValidateData(data model.PluginInfo) error {
	for _, event := range data.Events {
		if strings.TrimSpace(event.CreatedAt) == "" {
			return ErrInvalidCreatedAtFiled
		}

		if strings.TrimSpace(event.Type) == "" {
			return ErrInvalidTypeFiled
		}
	}

	return nil
}
