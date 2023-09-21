package dispatcher

func NewService(sender Sender, storage Storage) *Service {
	return &Service{sender: sender, storage: storage}
}

type Service struct {
	sender  Sender
	storage Storage
}

func (s *Service) SendData() error {
	filesToSend, err := s.storage.GetFilesToSend()
	if err != nil {
		return err
	}

	for _, file := range filesToSend {
		dataModels, err := s.storage.ReadDataToSend(file)

		if err != nil {
			return err
		}

		if err := s.sender.Send(dataModels); err != nil {
			return err
		}

		if err := s.storage.ClearSentData(file); err != nil {
			return err
		}
	}

	return nil
}
