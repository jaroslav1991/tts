package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
)

var urls = []string{
	"http://tts-tts-service-1:9191",
	"http://tts-tts-service-2:9191",
	"http://tts-tts-service-3:9191",
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// правильнее сделать proxy
		target := urls[rand.Intn(len(urls))]
		r, err := http.DefaultClient.Get(target + request.URL.Path)
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			b, _ := io.ReadAll(r.Body)
			writer.Write(b)
		}
	})

	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatalln(err)
	}
}
