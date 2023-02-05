package data

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var now = time.Date(2007, 02, 05, 16, 31, 17, 0, time.UTC)

func freezeTime(t *testing.T) func() {
	t.Helper()

	currentTime = func() time.Time {
		return now
	}

	unFreezeTime := func() {
		currentTime = time.Now
	}

	return unFreezeTime
}

func TestStorage_FixDataToSend_Positive(t *testing.T) {
	unFreezeTime := freezeTime(t)
	defer unFreezeTime()

	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	fixData := Storage{NewStatsFileName: f.Name()}

	actualNewName, actualError := fixData.FixDataToSend()
	expectedNewName := fmt.Sprintf("%s.%d", f.Name(), currentTime().UnixNano())

	assert.NoError(t, actualError)
	assert.NoError(t, os.Remove(actualNewName))
	assert.Equal(t, expectedNewName, actualNewName)
}

func TestStorage_FixDataToSend_Negative(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())
	assert.NoError(t, os.Remove(f.Name()))

	fixData := Storage{NewStatsFileName: f.Name()}

	actualNewName, actualError := fixData.FixDataToSend()
	assert.Empty(t, actualNewName)
	assert.ErrorIs(t, actualError, os.ErrNotExist)
}

func TestStorage_ClearSentData_Negative(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())
	assert.NoError(t, os.Remove(f.Name()))

	fixData := Storage{NewStatsFileName: ""}

	assert.ErrorIs(t, fixData.ClearSentData(f.Name()), os.ErrNotExist)
}

func TestStorage_ClearSentData_Positive(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	fixData := Storage{NewStatsFileName: ""}

	assert.NoError(t, fixData.ClearSentData(f.Name()))
}
