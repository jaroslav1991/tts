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
		Program:     "",
		Duration:    0,
		PathProject: "",
	}

	aggregatedData := model.AggregatorInfo{CurrentGitBranch: ""}

	preparedData := []byte(`prepared data`)

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(data).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(data, aggregatedData).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(nil)

	service := NewService(reader, validator, aggregator, preparer, saver)
	assert.NoError(t, service.SaveData(request))
}

func TestService_SaveData_Negative_SaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	aggregatedData := model.AggregatorInfo{CurrentGitBranch: ""}

	data := model.PluginInfo{
		Program:     "",
		Duration:    0,
		PathProject: "",
	}

	preparedData := []byte(`prepared data`)

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(data).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(data, aggregatedData).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(err)

	///
	service := NewService(reader, validator, aggregator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_PrepareError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:     "",
		Duration:    0,
		PathProject: "",
	}
	aggregatedData := model.AggregatorInfo{CurrentGitBranch: ""}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(data).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(data, aggregatedData).Return(nil, err)

	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_AggregatorError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:     "",
		Duration:    0,
		PathProject: "",
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(data).Return(model.AggregatorInfo{}, err)

	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_ValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	data := model.PluginInfo{
		Program:     "",
		Duration:    0,
		PathProject: "",
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(data, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(data).Return(err)

	aggregator := NewMockDataAggregator(ctrl)
	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)
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
	aggregator := NewMockDataAggregator(ctrl)
	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)
	assert.Error(t, service.SaveData(request))
}
