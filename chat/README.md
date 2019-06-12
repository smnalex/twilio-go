# Twilio Programmable Chat

Client for [Twilio Programmable Chat](https://www.twilio.com/docs/chat) API.

## Documentation
[GoDoc](https://godoc.org/github.com/smnalex/twilio-go/chat)

## Usage

### Services
```go
import (
    "github.com/twilio-go/"
    twchat "github.com/twilio-go/chat"
)
func main() {
    // Configuration required for twilio services
    configuration := twilio.NewContext()

    // Chat client
    chat, err := twchat.New(configuration)
    if err != nil {
        log.Fatal(err)
    }
    ...
}
```
