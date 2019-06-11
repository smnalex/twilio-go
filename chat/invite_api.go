package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smnalex/twilio-go"
)

type inviteAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Channels/{Channel SID}/Invites/{Invite SID}
// https://www.twilio.com/docs/chat/rest/invites#read-an-invite-to-a-channel
func (api inviteAPI) Read(ctx context.Context, serviceSID, channelSID, inviteSID string) (Invite, error) {
	var inv Invite
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Invites/%s", serviceSID, channelSID, inviteSID))
	if err != nil {
		return inv, err
	}
	err = json.Unmarshal(data, &inv)
	return inv, err
}

// POST /Services/{Service SID}/Channels/{Channel SID}/Invites
// https://www.twilio.com/docs/chat/rest/invites#create-an-invite-to-a-channel
func (api inviteAPI) Create(ctx context.Context, serviceSID, channelSID string, body InviteCreateParams) (Invite, error) {
	var inv Invite
	data, err := api.client.Post(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Invites", serviceSID, channelSID), body.encode())
	if err != nil {
		return inv, err
	}
	err = json.Unmarshal(data, &inv)
	return inv, err
}

// DELETE /Services/{Service SID}/Channels/{Channel SID}/Invites/{Invite SID}
// https://www.twilio.com/docs/chat/rest/invites#delete-an-invite-from-a-channel
func (api inviteAPI) Delete(ctx context.Context, serviceSID, channelSID, inviteSID string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Channels/%s/Invites/%s", serviceSID, channelSID, inviteSID))
	return err
}
