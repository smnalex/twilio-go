package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type roleAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Roles/{Role SID}
// https://www.twilio.com/docs/chat/rest/roles#retrieve-a-role
func (r roleAPI) Read(ctx context.Context, serviceSid, roleSid string) (Role, error) {
	var role Role
	data, err := r.client.Get(ctx, fmt.Sprintf("/Services/%s/Roles/%s", serviceSid, roleSid))
	if err != nil {
		return role, err
	}
	err = json.Unmarshal(data, &role)
	return role, err
}

// POST /Services/{Service SID}/Roles
// https://www.twilio.com/docs/chat/rest/roles#create-a-role
func (r roleAPI) Create(ctx context.Context, serviceSid string, body RoleCreateParams) (Role, error) {
	return r.post(ctx, fmt.Sprintf("/Services/%s/Roles", serviceSid), body.encode())
}

// POST /Services/{Service SID}/Roles/{Role SID}
// https://www.twilio.com/docs/chat/rest/roles#update-a-role
func (r roleAPI) Update(ctx context.Context, serviceSid, roleSid string, body RoleUpdateParams) (Role, error) {
	return r.post(ctx, fmt.Sprintf("/Services/%s/Roles/%s", serviceSid, roleSid), body.encode())
}

// DELETE /Services/{Service SID}/Roles/{Role SID}
// https://www.twilio.com/docs/chat/rest/roles#delete-a-role
func (r roleAPI) Delete(ctx context.Context, serviceSid, roleSid string) error {
	_, err := r.client.Delete(ctx, fmt.Sprintf("/Services/%s/Roles/%s", serviceSid, roleSid))
	return err
}

func (r roleAPI) post(ctx context.Context, path string, body io.Reader) (Role, error) {
	var role Role
	data, err := r.client.Post(ctx, path, body)
	if err != nil {
		return role, err
	}
	err = json.Unmarshal(data, &role)
	return role, err
}
