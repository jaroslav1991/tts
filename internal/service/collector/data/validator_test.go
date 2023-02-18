package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestValidator_ValidateData_Positive(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Program:  "testPro",
		Duration: 5,
	})
	assert.NoError(t, actualErr)
}

func TestValidator_ValidateData_Negative_Program_Field(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Program:  " ",
		Duration: 5,
	})
	assert.ErrorIs(t, actualErr, ErrInvalidProgramField)
}

func TestValidator_ValidateData_Negative_Duration_Field(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.PluginInfo{
		Program:  "test",
		Duration: 0,
	})
	assert.ErrorIs(t, actualErr, ErrInvalidDurationField)
}
