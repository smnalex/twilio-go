package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smnalex/twilio-go"
)

type userChannelAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Instance SID}/Users/{User SID}/Channels
// https://www.twilio.com/docs/chat/rest/user-channels#list-all-user-channels
func (api userChannelAPI) List(ctx context.Context, serviceSid, userSid string) (UserChannelList, error) {
	var chanList UserChannelList
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Users/%s/Channels", serviceSid, userSid))
	if err != nil {
		return chanList, err
	}
	err = json.Unmarshal(data, &chanList)
	return chanList, err
}
