package http

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := httptest.NewRequest(http.MethodPost, "/", nil)

	service := NewMockService(ctrl)
	service.EXPECT().SaveData(request).Return(nil)

	handler := NewHandler(service)

	responseWriter := httptest.NewRecorder()
	handler(responseWriter, request)

	assert.Equal(t, http.StatusOK, responseWriter.Code)

	body, err := io.ReadAll(responseWriter.Body)
	assert.NoError(t, err)
	assert.Empty(t, body)
}

func TestHandler_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := httptest.NewRequest(http.MethodPost, "/", nil)

	service := NewMockService(ctrl)
	service.EXPECT().SaveData(request).Return(errors.New("service error"))

	handler := NewHandler(service)

	responseWriter := httptest.NewRecorder()
	handler(responseWriter, request)

	assert.Equal(t, http.StatusInternalServerError, responseWriter.Code)

	body, err := io.ReadAll(responseWriter.Body)
	assert.NoError(t, err)
	assert.Equal(t, `service error`, string(body))
}
