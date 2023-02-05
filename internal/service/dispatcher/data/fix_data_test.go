package data

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFixData_FixDataToSend(t *testing.T) {

	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)

	f.Close()

	fixData := FixData{fileName: f.Name()}

	newName, actualError := fixData.FixDataToSend()
	assert.NoError(t, actualError)
	assert.NoError(t, os.Remove(newName))
}
