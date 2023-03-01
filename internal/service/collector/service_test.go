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
		CliType:       "macos",
		CliVersion:    "3.9.1",
		DeviceName:    "bambook",
		Events: []model.Events{
			{
				CreatedAt: "2022-02-02 10:00:00",
				Type:      "some-type",
				Project:   "some project",
				Language:  "golang",
				Target:    "some target",
				Params: map[string]any{
					"param1": "value1",
				},
			},
		},
	}

	aggregatedData := model.AggregatorInfo{
		GitBranchesByEventUID: map[string]string{
			"some-uuid": "some-branch",
		},
	}

	preparedData := []byte(`prepared dataFromPlugin`)

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	dataFromPluginWithUUID := model.PluginInfo{
		PluginType:    "type",
		PluginVersion: "1.0.0",
		CliType:       "macos",
		CliVersion:    "3.9.1",
		DeviceName:    "bambook",
		Events: []model.Events{
			{
				Uid:       "some-uuid",
				CreatedAt: "2022-02-02 10:00:00",
				Type:      "some-type",
				Project:   "some project",
				Language:  "golang",
				Target:    "some target",
				Params: map[string]any{
					"param1": "value1",
				},
			},
		},
	}

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithUUID).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(dataFromPluginWithUUID, aggregatedData).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(nil)

	service := NewService(reader, validator, aggregator, preparer, saver)
	service.uudGenFn = func() string {
		return "some-uuid"
	}

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

	dataFromPluginWithUUID := model.PluginInfo{
		Events: []model.Events{
			{
				Uid: "some-uuid",
			},
		},
	}

	aggregatedData := model.AggregatorInfo{
		GitBranchesByEventUID: map[string]string{
			"some-uuid": "some-branch",
		},
	}

	preparedData := []byte(`prepared data`)

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithUUID).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(dataFromPluginWithUUID, aggregatedData).Return(preparedData, nil)

	saver := NewMockDataSaver(ctrl)
	saver.EXPECT().SaveData(preparedData).Return(err)

	service := NewService(reader, validator, aggregator, preparer, saver)
	service.uudGenFn = func() string {
		return "some-uuid"
	}

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

	dataFromPluginWithUUID := model.PluginInfo{
		Events: []model.Events{
			{
				Uid: "some-uuid",
			},
		},
	}

	aggregatedData := model.AggregatorInfo{
		GitBranchesByEventUID: map[string]string{
			"some-uuid": "some-branch",
		},
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithUUID).Return(aggregatedData, nil)

	preparer := NewMockDataPreparer(ctrl)
	preparer.EXPECT().PrepareData(dataFromPluginWithUUID, aggregatedData).Return(nil, err)

	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)
	service.uudGenFn = func() string {
		return "some-uuid"
	}
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

	dataFromPluginWithUUID := model.PluginInfo{
		Events: []model.Events{
			{
				Uid: "some-uuid",
			},
		},
	}

	err := errors.New("some error")

	reader := NewMockDataReader(ctrl)
	reader.EXPECT().ReadData("some request").Return(dataFromPlugin, nil)

	validator := NewMockDataValidator(ctrl)
	validator.EXPECT().ValidateData(dataFromPlugin).Return(nil)

	aggregator := NewMockDataAggregator(ctrl)
	aggregator.EXPECT().Aggregate(dataFromPluginWithUUID).Return(model.AggregatorInfo{}, err)

	preparer := NewMockDataPreparer(ctrl)
	saver := NewMockDataSaver(ctrl)

	service := NewService(reader, validator, aggregator, preparer, saver)
	service.uudGenFn = func() string {
		return "some-uuid"
	}
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
	service.uudGenFn = func() string {
		return "some-uuid"
	}
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
	service.uudGenFn = func() string {
		return "some-uuid"
	}
	assert.Error(t, service.SaveData(request))
}
