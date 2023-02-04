package data

import (
	"testing"

	"github.com/jaroslav1991/tts/internal/service/model"
	"github.com/stretchr/testify/assert"
)

func TestValidator_ValidateData_Positive(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.DataModel{
		Program:  "testPro",
		Duration: 5,
	})
	assert.NoError(t, actualErr)
}

func TestValidator_ValidateData_Negative_Program_Field(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.DataModel{
		Program:  " ",
		Duration: 5,
	})
	assert.ErrorIs(t, actualErr, ErrInvalidProgramField)
}

func TestValidator_ValidateData_Negative_Duration_Field(t *testing.T) {
	validator := Validator{}

	actualErr := validator.ValidateData(model.DataModel{
		Program:  "test",
		Duration: 0,
	})
	assert.ErrorIs(t, actualErr, ErrInvalidDurationField)
}
