package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestValidator_ValidateData_Positive(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.NoError(t, actualErr)
}

func TestValidator_ValidateData_Negative_PluginType_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidPluginTypeField)
}

func TestValidator_ValidateData_Negative_PluginVersion_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidPluginVersionField)
}

func TestValidator_ValidateData_Negative_CliType_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidCliTypeFiled)
}

func TestValidator_ValidateData_Negative_CliVersion_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidCliVersionFiled)
}

func TestValidator_ValidateData_Negative_Uid_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidUidFiled)
}

func TestValidator_ValidateData_Negative_CreatedAt_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidCreatedAtFiled)
}

func TestValidator_ValidateData_Negative_Type_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidTypeFiled)
}
