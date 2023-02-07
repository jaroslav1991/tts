package data

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSender_Send_Positive(t *testing.T) {
	actualData := []model.DataModel{{Program: "test1", Duration: 2}}

	sender := Sender{}
	actualErr := sender.Send(actualData)
	assert.NoError(t, actualErr)
}
