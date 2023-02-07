package data

import (
	"encoding/json"
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"os"
	"strings"
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

	fixData := Storage{NewStatsFileName: f.Name(), FilePath: os.TempDir()}

	actualNewName, actualError := fixData.FixDataToSend()
	expectedNewName := fmt.Sprintf("%s%d", fixData.FilePath+string(os.PathSeparator), currentTime().UnixNano())

	assert.NoError(t, actualError)
	assert.NoError(t, os.Remove(actualNewName))
	assert.Equal(t, expectedNewName, actualNewName)
}

func TestStorage_FixDataToSend_Negative(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "testfile")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())
	assert.NoError(t, os.Remove(f.Name()))

	fixData := Storage{NewStatsFileName: f.Name(), FilePath: os.TempDir()}

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

func TestStorage_GetFilesToSend_Positive(t *testing.T) {
	tempDir := os.TempDir() + string(os.PathSeparator) + fmt.Sprintf("test%d", time.Now().UnixNano())

	err := os.Mkdir(tempDir, os.ModePerm)
	assert.NoError(t, err)

	defer os.RemoveAll(tempDir)

	f, err := os.CreateTemp(tempDir, "testfile1")
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	f2, err := os.CreateTemp(tempDir, "testfile2")
	assert.NoError(t, err)
	assert.NoError(t, f2.Close())

	baseNameF1 := strings.ReplaceAll(f.Name(), tempDir+string(os.PathSeparator), "")
	baseNameF2 := strings.ReplaceAll(f2.Name(), tempDir+string(os.PathSeparator), "")

	storage := Storage{FilePath: tempDir}
	expectedRes := []string{baseNameF1, baseNameF2}

	actualRes, actErr := storage.GetFilesToSend()
	assert.NoError(t, actErr)
	assert.Equal(t, expectedRes, actualRes)
}

func TestStorage_GetFilesToSend_Negative(t *testing.T) {
	tempDir := os.TempDir() + string(os.PathSeparator) + fmt.Sprintf("test%d", time.Now().UnixNano())

	storage := Storage{FilePath: tempDir}

	_, actErr := storage.GetFilesToSend()
	assert.Error(t, actErr)

}

func TestStorage_ReadDataToSend_Positive(t *testing.T) {
	var models []model.DataModel

	expectedModel := []model.DataModel{{Program: "test1", Duration: 2}, {Program: "test2", Duration: 3}}
	expectedModelBytes := []byte(`[{"Program": "test1", "Duration": 2},{"Program": "test2", "Duration": 3}]`)

	err := json.Unmarshal(expectedModelBytes, &models)
	assert.NoError(t, err)

	tempDir := os.TempDir() + string(os.PathSeparator) + fmt.Sprintf("%d", time.Now().UnixNano())

	err = os.Mkdir(tempDir, os.ModePerm)
	assert.NoError(t, err)

	defer os.RemoveAll(tempDir)

	file, err := os.CreateTemp(tempDir, "testingFile")
	assert.NoError(t, err)

	_, err = file.Write(expectedModelBytes)
	assert.NoError(t, err)
	assert.NoError(t, file.Close())

	storage := Storage{FilePath: tempDir}

	actualData, actErr := storage.ReadDataToSend(file.Name())
	assert.NoError(t, actErr)
	assert.Equal(t, actualData, expectedModel)
}

func TestStorage_ReadDataToSend_UnmarshalErr(t *testing.T) {
	tempDir := os.TempDir() + string(os.PathSeparator) + fmt.Sprintf("%d", time.Now().UnixNano())

	err := os.Mkdir(tempDir, os.ModePerm)
	assert.NoError(t, err)

	defer os.RemoveAll(tempDir)

	file, err := os.CreateTemp(tempDir, "testingFile")
	assert.NoError(t, err)
	assert.NoError(t, file.Close())

	storage := Storage{FilePath: tempDir}

	_, actErr := storage.ReadDataToSend(file.Name())
	assert.ErrorIs(t, actErr, ErrUnmarshalData)
}

func TestStorage_ReadDataToSend_ReadFileErr(t *testing.T) {
	tempDir := os.TempDir() + string(os.PathSeparator) + fmt.Sprintf("%d", time.Now().UnixNano())

	err := os.Mkdir(tempDir, os.ModePerm)
	assert.NoError(t, err)

	defer os.RemoveAll(tempDir)

	file, err := os.CreateTemp(tempDir, "testingFile")
	assert.NoError(t, err)
	assert.NoError(t, file.Close())
	assert.NoError(t, os.Remove(file.Name()))

	storage := Storage{FilePath: tempDir}

	_, actErr := storage.ReadDataToSend(file.Name())
	assert.Error(t, actErr)
	assert.ErrorIs(t, actErr, os.ErrNotExist)
}
