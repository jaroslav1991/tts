package data

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSaver_SaveData(t *testing.T) {
	err := os.Mkdir("stats", os.ModePerm)
	assert.NoError(t, err)

	saver := Saver{NewStatsFileName: "stats", AuthKey: "123"}

	actualErr := saver.SaveData([]byte(`test`))
	assert.NoError(t, actualErr)

	actualData, err := os.ReadFile("stats/123")
	assert.NoError(t, err)
	assert.Equal(t, "test\n", string(actualData))

	os.Remove("stats/123")
	os.Remove("stats")
}

func TestSaver_SaveData_InvalidFileName(t *testing.T) {
	err := os.Mkdir("stats", os.ModePerm)
	assert.NoError(t, err)

	saver := Saver{NewStatsFileName: "stats", AuthKey: ""}

	actualErr := saver.SaveData([]byte(`test`))
	assert.ErrorIs(t, actualErr, ErrCantOpenFile)
	os.Remove("stats")
}

func TestSaver_SaveData_InvalidPath(t *testing.T) {
	saver := Saver{NewStatsFileName: "", AuthKey: "123"}

	actualErr := saver.SaveData([]byte(`test`))
	assert.ErrorIs(t, actualErr, ErrCantCreatePath)
}
