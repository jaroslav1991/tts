package sender

//func TestSender_Send_Positive(t *testing.T) {
//	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		body, err := io.ReadAll(request.Body)
//		assert.NoError(t, err)
//		assert.Equal(t, `[{"pluginType":"1","pluginVersion":"1","cliType":"1","cliVersion":"1","deviceName":"1","Events":{"uid":"1","createdAt":"1","type":"1","project":"1","language":"1","target":"1","branch":"master"}}]`, string(body))
//	}))
//
//	actualData := []model.DataModel{
//		{
//			PluginInfo: model.PluginInfo{
//				PluginType:    "1",
//				PluginVersion: "1",
//				CliType:       "1",
//				CliVersion:    "1",
//				DeviceName:    "1",
//				Events: model.Events{
//					Uid:       "1",
//					CreatedAt: "1",
//					Type:      "1",
//					Project:   "1",
//					Language:  "1",
//					Target:    "1",
//					Branch:    "master",
//					Params:    nil,
//				},
//			},
//			AggregatorInfo: model.AggregatorInfo{
//				CurrentGitBranch: "master",
//			},
//		},
//	}
//
//	sender := Sender{HttpAddr: server.URL}
//	actualErr := sender.Send(actualData)
//	assert.NoError(t, actualErr)
//}
