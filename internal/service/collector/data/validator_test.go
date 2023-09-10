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
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
		OSName:        "",
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
				Timezone:       "",
				Params:         nil,
			},
		},
	})
	assert.NoError(t, actualErr)
}

func TestValidator_ValidateData_Negative_CliType_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "",
		PluginVersion: "1",
		CliVersion:    "1.1.0",
		OSName:        "",
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
	assert.ErrorIs(t, actualErr, ErrInvalidCliTypeField)
}

func TestValidator_ValidateData_Negative_CliVersion_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "",
		PluginVersion: "1",
		CliType:       "windowsOS",
		OSName:        "",
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
	assert.ErrorIs(t, actualErr, ErrInvalidCliVersionField)
}

func TestValidator_ValidateData_Negative_PluginVersion_Filed(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "",
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
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
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
		OSName:        "",
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
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
		OSName:        "",
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
