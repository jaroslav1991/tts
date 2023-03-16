package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestValidator_ValidateData_Positive(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	})
	assert.NoError(t, actualErr)
}

func TestValidator_ValidateData_Negative_Uid_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Uid:           "",
		PluginType:    "1",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidUidField)
}

func TestValidator_ValidateData_Negative_PluginType_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidPluginTypeField)
}

func TestValidator_ValidateData_Negative_PluginVersion_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidPluginVersionField)
}

func TestValidator_ValidateData_Negative_CreatedAt_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidCreatedAtFiled)
}

func TestValidator_ValidateData_Negative_Type_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	})
	assert.ErrorIs(t, actualErr, ErrInvalidTypeFiled)
}
