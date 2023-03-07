package data

import (
	"errors"
	"log"
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
		log.Println(ErrInvalidPluginTypeField)
		return ErrInvalidPluginTypeField
	}

	if strings.TrimSpace(data.PluginVersion) == "" {
		log.Println(ErrInvalidPluginVersionField)
		return ErrInvalidPluginVersionField
	}

	if strings.TrimSpace(data.CliType) == "" {
		log.Println(ErrInvalidCliTypeFiled)
		return ErrInvalidCliTypeFiled
	}

	if strings.TrimSpace(data.CliVersion) == "" {
		log.Println(ErrInvalidCliVersionFiled)
		return ErrInvalidCliVersionFiled
	}

	for _, event := range data.Events {
		if strings.TrimSpace(event.CreatedAt) == "" {
			log.Println(ErrInvalidCreatedAtFiled)
			return ErrInvalidCreatedAtFiled
		}

		if strings.TrimSpace(event.Type) == "" {
			log.Println(ErrInvalidTypeFiled)
			return ErrInvalidTypeFiled
		}
	}

	return nil
}
