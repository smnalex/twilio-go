package chat

import (
	"fmt"
	"os"

	"github.com/smnalex/twilio-go"
)

// Chat programmable chat interface
type Chat struct {
	Bindings     interface{}
	Channels     interface{}
	Credentials  interface{}
	Members      interface{}
	Invites      interface{}
	Messages     interface{}
	Roles        interface{}
	Services     ServiceResource
	Users        interface{}
	UserChannels interface{}
	Media        interface{}
}

func chatEndpointForRegion(region string) string {
	url := os.Getenv("TWILIO_CHAT_HOST")
	if url == "" && region != "" {
		return fmt.Sprintf("https://chat.%s.twilio.com/v2", region)
	} else if url == "" {
		return "https://chat.twilio.com/v2"
	}
	return url
}

// New returns a chat instance with a base url set to `https://chat.twilio.com/v2`
// if `TWILIO_CHAT_HOST` env not defined.
func New(tctx twilio.TwilioContext) (*Chat, error) {
	var chat Chat

	client, err := twilio.NewHTTPClient(
		tctx.AccountSID,
		tctx.AuthToken,
		chatEndpointForRegion(tctx.Region),
		tctx.HTTPClient,
	)
	if err != nil {
		return nil, err
	}

	{
		chat.Services = ServiceResource{serviceAPI{client}}
	}

	return &chat, nil
}
