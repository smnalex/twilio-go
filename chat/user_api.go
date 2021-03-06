package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type userAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Users/{Identity}
// GET /Services/{Service SID}/Users/{User SID}
// https://www.twilio.com/docs/chat/rest/users#retrieve-a-user
func (api userAPI) Read(ctx context.Context, serviceSid, identity string) (User, error) {
	var usr User
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Users/%s", serviceSid, identity))
	if err != nil {
		return usr, err
	}
	err = json.Unmarshal(data, &usr)
	return usr, err
}

// POST /Services/{Service SID}/Users
// https://www.twilio.com/docs/chat/rest/users#create-a-user
func (api userAPI) Create(ctx context.Context, serviceSid string, body UserCreateParams) (User, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Users", serviceSid), body.encode())
}

// POST /Services/{Service SID}/Users/{User SID}
// https://www.twilio.com/docs/chat/rest/users#update-a-user
func (api userAPI) Update(ctx context.Context, serviceSid, identity string, body UserUpdateParams) (User, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s/Users/%s", serviceSid, identity), body.encode())
}

// DELETE /Services/{Service SID}/Users/{User SID}
// https://www.twilio.com/docs/chat/rest/users#delete-a-user
func (api userAPI) Delete(ctx context.Context, serviceSid, identity string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Users/%s", serviceSid, identity))
	return err
}

func (api userAPI) post(ctx context.Context, path string, body io.Reader) (User, error) {
	var usr User
	data, err := api.client.Post(ctx, path, body)
	if err != nil {
		return usr, err
	}
	err = json.Unmarshal(data, &usr)
	return usr, err
}
