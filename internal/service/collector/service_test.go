package collector

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestService_SaveData_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:  "",
		Duration: 0,
	}

	preparedData := []byte(`prepared data`)

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(data).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(nil)

	service := NewService(reader, validator, preparer, saver)
	assert.NoError(t, service.SaveData(request))
}

func TestService_SaveData_Negative_SaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:  "",
		Duration: 0,
	}

	preparedData := []byte(`prepared data`)

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(data).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(err)

	///
	service := NewService(reader, validator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_PrepareError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:  "",
		Duration: 0,
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(data).Return(nil, err)

	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_ValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:  "",
		Duration: 0,
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(err)

	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_ReadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(model.PluginInfo{}, err)

	validator := NewMockDataValidator(ctrl)
	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}
