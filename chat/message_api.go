package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type messageAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Channels/{Channel SID}/Messages/{Message SID}
// https://www.twilio.com/docs/chat/rest/messages#retrieve-a-message-from-a-channel
func (api messageAPI) Read(ctx context.Context, serviceSid, channelSid, messageSid string) (Message, error) {
	var msg Message
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Messages/%s", serviceSid, channelSid, messageSid))
	if err != nil {
		return msg, err
	}
	err = json.Unmarshal(data, &msg)
	return msg, err
}

// POST /Services/{Service SID}/Channels/{Channel SID}/Messages
// https://www.twilio.com/docs/chat/rest/messages#send-a-message-to-a-channel
func (api messageAPI) Send(ctx context.Context, serviceSid, channelSid string, body MessageCreateParams) (Message, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Messages", serviceSid, channelSid), body.encode())
}

// POST /Services/{Service SID}/Channels/{Channel SID}/Messages/{Message SID}
// https://www.twilio.com/docs/chat/rest/messages#update-an-existing-message
func (api messageAPI) Update(ctx context.Context, serviceSid, channelSid, messageSid string, body MessageUpdateParams) (Message, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Messages/%s", serviceSid, channelSid, messageSid), body.encode())
}

func (api messageAPI) Delete(ctx context.Context, serviceSid, channelSid, messageSid string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Messages/%s", serviceSid, channelSid, messageSid))
	return err
}

func (api messageAPI) post(ctx context.Context, path string, body io.Reader) (Message, error) {
	var msg Message
	data, err := api.client.Post(ctx, path, body)
	if err != nil {
		return msg, err
	}
	err = json.Unmarshal(data, &msg)
	return msg, err
}
