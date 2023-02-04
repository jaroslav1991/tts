package data

import (
	"github.com/jaroslav1991/tts/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_ValidateData_Positive(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(service.DataModel{
		Program:  "testPro",
		Duration: 5,
	})
	assert.NoError(t, actualErr)
}

func TestValidator_ValidateData_Negative_Program_Field(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(service.DataModel{
		Program:  " ",
		Duration: 5,
	})
	assert.ErrorIs(t, actualErr, ErrInvalidProgramField)
}

func TestValidator_ValidateData_Negative_Duration_Field(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(service.DataModel{
		Program:  "test",
		Duration: 0,
	})
	assert.ErrorIs(t, actualErr, ErrInvalidDurationField)
}
