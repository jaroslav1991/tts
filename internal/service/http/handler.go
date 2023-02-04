package http

import "net/http"

func NewHandler(service Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := service.SaveData(request); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}
