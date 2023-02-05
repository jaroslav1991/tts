package data

import (
	"fmt"
	"github.com/jaroslav1991/tts/internal/service/dispatcher"
	"os"
	"time"
)

type FixData struct {
	dispatcher.Storage
	fileName string
}

func (f *FixData) FixDataToSend() (string, error) {
	now := time.Now().UnixNano()
	newPath := fmt.Sprintf("%s.%d", f.fileName, now)

	return newPath, os.Rename(f.fileName, newPath)
}
