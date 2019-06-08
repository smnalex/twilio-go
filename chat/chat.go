package chat

import (
	"fmt"
	"os"

	"github.com/smnalex/twilio-go"
)

// Chat programmable chat interface
type Chat struct {
	Bindings     BindingResource
	Channels     ChannelResource
	Credentials  interface{}
	Members      interface{}
	Invites      interface{}
	Messages     interface{}
	Roles        RoleResource
	Services     ServiceResource
	Users        UserResource
	UserChannels UserChannelResource
}

// New returns a chat instance with a base url set to `https://chat.twilio.com/v2`
// if `TWILIO_CHAT_HOST` not set.
func New(tctx twilio.Context) (Chat, error) {
	var chatClient Chat

	client, err := twilio.NewHTTPClient(
		tctx.APIKey,
		tctx.APISecret,
		chatEndpointForRegion(tctx.Region),
		tctx.RequestHandler,
	)
	if err != nil {
		return chatClient, err
	}

	{
		chatClient.Roles = RoleResource{roleAPI{client}}
		chatClient.Services = ServiceResource{serviceAPI{client}}
	}
	return chatClient, nil
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
