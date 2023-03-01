package dispatcher

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestService_SendData_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "1",
				CliVersion:    "1",
				DeviceName:    "1",
				Events: model.Events{
					Uid:       "1",
					CreatedAt: "1",
					Type:      "1",
					Project:   "1",
					Language:  "1",
					Target:    "1",
					Branch:    "1",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "1"},
		},
	}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return("", nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file).Return(dataToSend, nil)

	sender := NewMockSender(ctrl)
	sender.EXPECT().Send(dataToSend).Return(nil)

	storage.EXPECT().ClearSentData(file).Return(nil)

	service := NewService(sender, storage)
	assert.NoError(t, service.SendData())
}

func TestService_SendData_Positive_WhenNoDataToFix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "1",
				CliVersion:    "1",
				DeviceName:    "1",
				Events: model.Events{
					Uid:       "1",
					CreatedAt: "1",
					Type:      "1",
					Project:   "1",
					Language:  "1",
					Target:    "1",
					Branch:    "1",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "1"},
		},
	}
	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return("", errors.New("no new data"))

	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file).Return(dataToSend, nil)

	sender := NewMockSender(ctrl)
	sender.EXPECT().Send(dataToSend).Return(nil)

	storage.EXPECT().ClearSentData(file).Return(nil)

	service := NewService(sender, storage)
	assert.NoError(t, service.SendData())
}

func TestService_SendData_Positive_MultiFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filesToSend := []string{"fileToSend1", "fileToSend2"}
	file1 := "fileToSend1"
	file2 := "fileToSend2"
	dataToSend := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "1",
				CliVersion:    "1",
				DeviceName:    "1",
				Events: model.Events{
					Uid:       "1",
					CreatedAt: "1",
					Type:      "1",
					Project:   "1",
					Language:  "1",
					Target:    "1",
					Branch:    "1",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "1"},
		},
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "2",
				PluginVersion: "2",
				CliType:       "2",
				CliVersion:    "2",
				DeviceName:    "2",
				Events: model.Events{
					Uid:       "2",
					CreatedAt: "2",
					Type:      "2",
					Project:   "2",
					Language:  "2",
					Target:    "2",
					Branch:    "2",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "2"},
		},
	}
	storage := NewMockStorage(ctrl)
	sender := NewMockSender(ctrl)

	storage.EXPECT().FixDataToSend().Return("", nil)
	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file1).Return(dataToSend, nil)
	sender.EXPECT().Send(dataToSend).Return(nil)
	storage.EXPECT().ClearSentData(file1).Return(nil)

	storage.EXPECT().ReadDataToSend(file2).Return(dataToSend, nil)
	sender.EXPECT().Send(dataToSend).Return(nil)
	storage.EXPECT().ClearSentData(file2).Return(nil)

	service := NewService(sender, storage)
	assert.NoError(t, service.SendData())
}

func TestService_SendData_Negative_ClearError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "1",
				CliVersion:    "1",
				DeviceName:    "1",
				Events: model.Events{
					Uid:       "1",
					CreatedAt: "1",
					Type:      "1",
					Project:   "1",
					Language:  "1",
					Target:    "1",
					Branch:    "1",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "1"},
		},
	}
	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return("", nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file).Return(dataToSend, nil)

	sender := NewMockSender(ctrl)
	sender.EXPECT().Send(dataToSend).Return(nil)

	storage.EXPECT().ClearSentData(file).Return(err)

	service := NewService(sender, storage)
	assert.Error(t, service.SendData())
}

func TestService_SendData_Negative_SenderError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "1",
				CliVersion:    "1",
				DeviceName:    "1",
				Events: model.Events{
					Uid:       "1",
					CreatedAt: "1",
					Type:      "1",
					Project:   "1",
					Language:  "1",
					Target:    "1",
					Branch:    "1",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "1"},
		},
	}
	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return("", nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file).Return(dataToSend, nil)

	sender := NewMockSender(ctrl)
	sender.EXPECT().Send(dataToSend).Return(err)

	service := NewService(sender, storage)
	assert.Error(t, service.SendData())
}

func TestService_SendData_Negative_ReadDataError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "1",
				CliVersion:    "1",
				DeviceName:    "1",
				Events: model.Events{
					Uid:       "1",
					CreatedAt: "1",
					Type:      "1",
					Project:   "1",
					Language:  "1",
					Target:    "1",
					Branch:    "1",
					Params:    nil,
				},
			},
			AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "testBranch"},
		},
	}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return("", nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file).Return(dataToSend, err)

	sender := NewMockSender(ctrl)

	service := NewService(sender, storage)
	assert.Error(t, service.SendData())
}

func TestService_SendData_Negative_GetFilesError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")

	filesToSend := []string{"fileToSend1"}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return("", nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, err)

	sender := NewMockSender(ctrl)

	service := NewService(sender, storage)
	assert.Error(t, service.SendData())
}
