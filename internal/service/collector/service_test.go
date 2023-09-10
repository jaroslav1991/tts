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

	dataFromPlugin := model.PluginInfo{
		PluginType:    "type",
		PluginVersion: "1.0.0",
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
		OSName:        "",
		IdeType:       "intellij ide",
		IdeVersion:    "2.1.1",
		Events: []model.Events{
			{
				Id:             "",
				CreatedAt:      "2022-02-02 10:00:00",
				Type:           "some-type",
				Project:        "some project",
				ProjectBaseDir: "some-base",
				Language:       "golang",
				Target:         "some target",
				Branch:         "",
				Timezone:       "123456789",
				Params: map[string]any{
					"param1": "value1",
				},
			},
		},
	}

	aggregatedData := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"some-base": "some-branch",
		},
		OSName: "windows",
		Id:     "7edc942b-da5c-4d33-9a25-f0c8b4bb4afc",
	}

	preparedData := []byte(`prepared dataFromPlugin`)

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPlugin).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(dataFromPlugin, aggregatedData).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(nil)

	service := NewService(reader, validator, aggregator, preparer, saver)

	assert.NoError(t, service.SaveData(request))
}

func TestService_SaveData_Negative_SaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	dataFromPlugin := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	dataFromPluginWithBranch := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	aggregatedData := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"some-base": "some-branch",
		},
		OSName: "windows",
		Id:     "7edc942b-da5c-4d33-9a25-f0c8b4bb4afc",
	}

	preparedData := []byte(`prepared data`)

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithBranch).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(dataFromPluginWithBranch, aggregatedData).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(err)

	service := NewService(reader, validator, aggregator, preparer, saver)

	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_PrepareError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	dataFromPlugin := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	dataFromPluginWithBranch := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	aggregatedData := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"some-base": "some-branch",
		},
		OSName: "windows",
		Id:     "7edc942b-da5c-4d33-9a25-f0c8b4bb4afc",
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithBranch).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(dataFromPluginWithBranch, aggregatedData).Return(nil, err)

	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)

	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_AggregatorError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	dataFromPlugin := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	dataFromPluginWithBranch := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithBranch).Return(model.AggregatorInfo{}, err)

	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)

	assert.Error(t, service.SaveData(request))
}

func TestService_SaveData_Negative_ValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := "some request"

	dataFromPlugin := model.PluginInfo{
		Events: []model.Events{
			{},
		},
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(err)

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
