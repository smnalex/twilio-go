package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smnalex/twilio-go"
)

type memberAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Channels/{Channel SID}/Members/{Member Identity}
// GET /Services/{Service SID}/Channels/{Channel SID}/Members/{Member SID}
// https://www.twilio.com/docs/chat/rest/members#retrieve-a-member-of-a-channel
func (api memberAPI) Read(ctx context.Context, serviceSid, channelSid, identity string) (Member, error) {
	var mem Member
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Members/%s", serviceSid, channelSid, identity))
	if err != nil {
		return mem, err
	}
	err = json.Unmarshal(data, &mem)
	return mem, err
}

// POST /Services/{Service SID}/Channels/{Channel SID}/Members
// https://www.twilio.com/docs/chat/rest/members#add-a-member-to-a-channel
func (api memberAPI) Add(ctx context.Context, serviceSid, channelSid string, body MemberCreateParams) (Member, error) {
	var mem Member
	data, err := api.client.Post(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Members", serviceSid, channelSid), body.encode())
	if err != nil {
		return mem, err
	}
	err = json.Unmarshal(data, &mem)
	return mem, err
}

// DELETE /Services/{Service SID}/Channels/{Channel SID}/Members/{Member SID}
// DELETE /Services/{Service SID}/Channels/{Channel SID}/Members/{Member Identity}
// https://www.twilio.com/docs/chat/rest/members#remove-a-member-from-a-channel
func (api memberAPI) Delete(ctx context.Context, serviceSid, channelSid, identity string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Members/%s", serviceSid, channelSid, identity))
	return err
}
