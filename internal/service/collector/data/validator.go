package data

import (
	"errors"
	"strings"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrInvalidCliTypeField       = errors.New("invalid cli type field")
	ErrInvalidCliVersionField    = errors.New("invalid cli version field")
	ErrInvalidPluginTypeField    = errors.New("invalid plugin type field")
	ErrInvalidPluginVersionField = errors.New("invalid plugin version field")
	ErrInvalidCreatedAtFiled     = errors.New("invalid created at field")
	ErrInvalidTypeFiled          = errors.New("invalid type field")
)

type Validator struct {
	collector.DataValidator
}

func (v Validator) ValidateData(data model.PluginInfo) error {
	if strings.TrimSpace(data.CliType) == "" {
		return ErrInvalidCliTypeField
	}

	if strings.TrimSpace(data.CliVersion) == "" {
		return ErrInvalidCliVersionField
	}

	if strings.TrimSpace(data.PluginType) == "" {
		return ErrInvalidPluginTypeField
	}

	if strings.TrimSpace(data.PluginVersion) == "" {
		return ErrInvalidPluginVersionField
	}

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
