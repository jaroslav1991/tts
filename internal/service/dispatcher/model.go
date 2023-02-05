package dispatcher

type DataToSend struct {
	Meta  DataToSendMeta   `json:"meta"`
	Stats []DataToSendStat `json:"stats"`
}

type DataToSendMeta struct {
	UserID int64
}

type DataToSendStat struct {
}
