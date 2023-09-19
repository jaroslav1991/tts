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

-s string                                                                     
Http address for sending events (default "http://localhost:8080/events")

-t string                                                                     
File for temporary storage of stats (default "./stats")  

## Example:
go run ./cmd/cli/main.go -d '{"events":[{"createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","projectBaseDir":"/mnt/c/Users/jaros/GolandProjects/tts","language":"golang","target":"C/Projects/Golang/cli-tts"},{"id":"qwerty","createdAt":"2022-02-1114:23:01","type":"modifyfile2","project":"someproject2","projectBaseDir":"/mnt/c/Users/jaros/GolandProjects/leetcode","language":"golang","target":"C/Projects/Golang/cli-tts"},{"createdAt":"2023-01-1114:23:01","type":"modifyfile3","project":"someproject3","projectBaseDir":"/mnt/c/Users/jaros/GolandProjects/tts","language":"golang","target":"C/Projects/Golang/cli-tts"}]}' -k 'b94f9d44-2cd9-4a30-aefc-1578a4eb9d6c'