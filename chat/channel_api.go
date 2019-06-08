package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type (
	// ChannelRepository interface for interacting with the channel chat api.
	ChannelRepository interface {
		Read(ctx context.Context, serviceSID, identity string) (Channel, error)
		Create(ctx context.Context, serviceSID string, body ChannelCreateParams) (Channel, error)
		Update(ctx context.Context, serviceSID, identity string, body ChannelUpdateParams) (Channel, error)
		Delete(ctx context.Context, serviceSID, identity string) error
	}

	channelAPI struct {
		client twilio.HTTPClient
	}
)

// GET /Services/{Service SID}/Channels/{Channel SID}
// GET /Services/{Service SID}/Channels/{Unique Name}
// https://www.twilio.com/docs/chat/rest/channels#retrieve-a-channel
func (api channelAPI) Read(ctx context.Context, sid, identity string) (Channel, error) {
	var chn Channel
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Channels/%s", sid, identity))
	if err != nil {
		return chn, err
	}
	err = json.Unmarshal(data, &chn)
	return chn, err
}

// POST /Services/{Service SID}/Channels/{Channel SID}
// POST /Services/{Service SID}/Channels/{Unique Name}
func (api channelAPI) Create(ctx context.Context, sid string, body ChannelCreateParams) (Channel, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Channels", sid), body.encode())
}

// POST /Services/{Service SID}/Channels/{Channel SID}
// POST /Services/{Service SID}/Channels/{Unique Name}
// https://www.twilio.com/docs/chat/rest/channels#update-a-channel
func (api channelAPI) Update(ctx context.Context, sid, identity string, body ChannelUpdateParams) (Channel, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Channels/%s", sid, identity), body.encode())
}

// DELETE /Services/{Service SID}/Channels/{Channel SID}
// DELETE /Services/{Service SID}/Channels/{Unique Name}
// https://www.twilio.com/docs/chat/rest/channels#delete-a-channel
func (api channelAPI) Delete(ctx context.Context, sid, identity string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Channels/%s", sid, identity))
	return err
}

func (api channelAPI) post(ctx context.Context, path string, body io.Reader) (Channel, error) {
	var chn Channel
	data, err := api.client.Post(ctx, path, body)
	if err != nil {
		return chn, err
	}
	err = json.Unmarshal(data, &chn)
	return chn, err
}
