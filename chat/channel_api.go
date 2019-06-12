package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type channelAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Channels/{Channel SID}
// GET /Services/{Service SID}/Channels/{Unique Name}
// https://www.twilio.com/docs/chat/rest/channels#retrieve-a-channel
func (api channelAPI) Read(ctx context.Context, serviceSid, identity string) (Channel, error) {
	var chn Channel
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Channels/%s", serviceSid, identity))
	if err != nil {
		return chn, err
	}
	err = json.Unmarshal(data, &chn)
	return chn, err
}

// POST /Services/{Service SID}/Channels/{Channel SID}
// POST /Services/{Service SID}/Channels/{Unique Name}
func (api channelAPI) Create(ctx context.Context, serviceSid string, body ChannelCreateParams) (Channel, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Channels", serviceSid), body.encode())
}

// POST /Services/{Service SID}/Channels/{Channel SID}
// POST /Services/{Service SID}/Channels/{Unique Name}
// https://www.twilio.com/docs/chat/rest/channels#update-a-channel
func (api channelAPI) Update(ctx context.Context, serviceSid, identity string, body ChannelUpdateParams) (Channel, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Channels/%s", serviceSid, identity), body.encode())
}

// DELETE /Services/{Service SID}/Channels/{Channel SID}
// DELETE /Services/{Service SID}/Channels/{Unique Name}
// https://www.twilio.com/docs/chat/rest/channels#delete-a-channel
func (api channelAPI) Delete(ctx context.Context, serviceSid, identity string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Channels/%s", serviceSid, identity))
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
