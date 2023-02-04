package dispatcher

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_SendData_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{{Program: "test1", Duration: 5}}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return(nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, nil)

	storage.EXPECT().ReadDataToSend(file).Return(dataToSend, nil)

	sender := NewMockSender(ctrl)
	sender.EXPECT().Send(dataToSend).Return(nil)

	storage.EXPECT().ClearSentData(file).Return(nil)

	service := NewService(sender, storage)
	assert.NoError(t, service.SendData())
}

func TestService_SendData_Negative_ClearError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")

	filesToSend := []string{"fileToSend1"}
	file := "fileToSend1"
	dataToSend := []model.DataModel{}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return(nil)

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
	dataToSend := []model.DataModel{}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return(nil)

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
	dataToSend := []model.DataModel{}

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return(nil)

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
	storage.EXPECT().FixDataToSend().Return(nil)

	storage.EXPECT().GetFilesToSend().Return(filesToSend, err)

	sender := NewMockSender(ctrl)

	service := NewService(sender, storage)
	assert.Error(t, service.SendData())
}

func TestService_SendData_Negative_FixDataError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")

	storage := NewMockStorage(ctrl)
	storage.EXPECT().FixDataToSend().Return(err)

	sender := NewMockSender(ctrl)

	service := NewService(sender, storage)
	assert.Error(t, service.SendData())
}
