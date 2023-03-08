package data

import (
	"errors"
	"strings"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector"
)

var (
	ErrInvalidPluginTypeField    = errors.New("invalid plugin type field")
	ErrInvalidPluginVersionField = errors.New("invalid plugin version field")
	ErrInvalidCliTypeFiled       = errors.New("invalid cli type field")
	ErrInvalidCliVersionFiled    = errors.New("invalid cli version field")
	ErrInvalidCreatedAtFiled     = errors.New("invalid created at field")
	ErrInvalidTypeFiled          = errors.New("invalid type field")
)

type Validator struct {
	collector.DataValidator
}

func (v Validator) ValidateData(data model.PluginInfo) error {
	if strings.TrimSpace(data.PluginType) == "" {
		return ErrInvalidPluginTypeField
	}

	if strings.TrimSpace(data.PluginVersion) == "" {
		return ErrInvalidPluginVersionField
	}

	if strings.TrimSpace(data.CliType) == "" {
		return ErrInvalidCliTypeFiled
	}

	if strings.TrimSpace(data.CliVersion) == "" {
		return ErrInvalidCliVersionFiled
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
