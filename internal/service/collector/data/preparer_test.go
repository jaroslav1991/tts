package data

import (
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestPreparer_PrepareData(t *testing.T) {
	preparer := Preparer{}

	data := model.DataModel{
		Program:  "test1",
		Duration: 0,
	}

	actualData, err := preparer.PrepareData(data)
	assert.NoError(t, err)
	assert.Equal(t, `{"Program":"test1","Duration":0}`, string(actualData))
}
