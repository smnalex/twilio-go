# Twilio-Go 

Clients for [Twilio](https://www.twilio.com/docs/) API.

## Install 

`go get github.com/smnalex/twilio-go`

## Documentation

[GoDoc](https://godoc.org/github.com/smnalex/twilio-go)

## Usage

### Programmable Chat
```go
import (
    "github.com/smnalex/twilio-go"
    "github.com/smnalex/twilio-go/chat"
)

func main() {
    // with http.DefaultClient
    configuration := twilio.New(accountSID, authToken, region)
    
    // with custom httpClient
    configuration := twilio.NewWithHTTP(accountSID, authToken, region, http.DefaultClient)

    chatClient, err := chat.New(configuration)
}
```

## Contirbutions
