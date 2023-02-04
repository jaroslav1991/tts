package data

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaver_SaveData(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)

	defer os.Remove(f.Name())
	defer f.Close()

	saver := Saver{FileName: f.Name()}

	actualErr := saver.SaveData([]byte(`test`))
	assert.NoError(t, actualErr)

	actualData, err := os.ReadFile(f.Name())
	assert.NoError(t, err)
	assert.Equal(t, "test\n", string(actualData))
}

func TestSaver_SaveData_InvalidFileName(t *testing.T) {
	saver := Saver{}

	actualErr := saver.SaveData([]byte(`test`))
	assert.ErrorIs(t, actualErr, ErrCantOpenFile)
}
