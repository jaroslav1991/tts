package data

import (
	"github.com/jaroslav1991/tts/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreparer_PrepareData(t *testing.T) {
	preparer := Preparer{}

	data := service.DataModel{
		Program:  "test1",
		Duration: 0,
	}

	actualData, err := preparer.PrepareData(data)
	assert.NoError(t, err)
	assert.Equal(t, `{"Program":"test1","Duration":0}`, string(actualData))
}
