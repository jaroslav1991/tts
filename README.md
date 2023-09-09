# tts

## How to debug locally

Start remote server mock:
```shell
make start-mock
```

Send test event through cli
```shell
make send-test-event
```

### How to send data

-d string                                                                     
Stats data in JSON format string  

-k string                                                                     
authorization key

-o string                                                                     
Path for temporary store files (default "./outbox")

-s string                                                                     
Http address for sending events (default "http://localhost:8080/events")

-t string                                                                     
File for temporary storage of stats (default "./stats")  

## Example:
go run ./cmd/cli/main.go -d '{"uid":"123e4567-e89b-12d3-a456-426655440000","pluginType":"jetbrains","pluginVersion":"1.0.0","ideType":"intellij idea","ideVersion":"2.1.1","events":[{"createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","projectBaseDir":"./","language":"golang","target":"C/Projects/Golang/cli-tts"}]}' -s "http://localhost:8181/events" 