package dispatcher

func NewService(sender Sender, storage Storage) *Service {
	return &Service{sender: sender, storage: storage}
}

type Service struct {
	sender  Sender
	storage Storage
}

func (s *Service) SendData() error {
	// rename file: temp -> temp_to_send_time
	if err := s.storage.FixDataToSend(); err != nil {
		return err
	}

	// get file list to send
	filesToSend, err := s.storage.GetFilesToSend()
	if err != nil {
		return err
	}

	for _, file := range filesToSend {
		// []model
		dataToSend, err := s.storage.ReadDataToSend(file)
		if err != nil {
			return err
		}

		if err := s.sender.Send(dataToSend); err != nil {
			return err
		}

		return s.storage.ClearSentData(file)
	}

	return nil
}
