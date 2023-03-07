package dispatcher

import (
	"log"
)

func NewService(sender Sender, storage Storage) *Service {
	return &Service{sender: sender, storage: storage}
}

type Service struct {
	sender  Sender
	storage Storage
}

func (s *Service) SendData() error {
	if _, err := s.storage.FixDataToSend(); err != nil {
		log.Println("non critical error: fix data to send: ", err)
	}

	filesToSend, err := s.storage.GetFilesToSend()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, file := range filesToSend {
		dataModels, err := s.storage.ReadDataToSend(file)

		if err != nil {
			log.Println(err)
			return err
		}

		if err := s.sender.Send(dataModels); err != nil {
			log.Println(err)
			return err
		}

		if err := s.storage.ClearSentData(file); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
