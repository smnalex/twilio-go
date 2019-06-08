package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type (
	// RoleRepository interface for interacting with the role chat api.
	RoleRepository interface {
		Read(ctx context.Context, serviceSID, roleSID string) (Role, error)
		Create(ctx context.Context, serviceSID string, body RoleCreateParams) (Role, error)
		Update(ctx context.Context, serviceSID, roleSID string, body RoleUpdateParams) (Role, error)
		Delete(ctx context.Context, serviceSID, roleSID string) error
	}

	roleAPI struct {
		client twilio.HTTPClient
	}
)

// GET /Services/{Service SID}/Roles/{Role SID}
// https://www.twilio.com/docs/chat/rest/roles#retrieve-a-role
func (r roleAPI) Read(ctx context.Context, sid, roleSID string) (Role, error) {
	var role Role
	data, err := r.client.Get(ctx, fmt.Sprintf("/Services/%s/Roles/%s", sid, roleSID))
	if err != nil {
		return role, err
	}
	err = json.Unmarshal(data, &role)
	return role, err
}

// POST /Services/{Service SID}/Roles
// https://www.twilio.com/docs/chat/rest/roles#create-a-role
func (r roleAPI) Create(ctx context.Context, sid string, body RoleCreateParams) (Role, error) {
	return r.post(ctx, fmt.Sprintf("/Services/%s/Roles", sid), body.encode())
}

// POST /Services/{Service SID}/Roles/{Role SID}
// https://www.twilio.com/docs/chat/rest/roles#update-a-role
func (r roleAPI) Update(ctx context.Context, sid, roleSID string, body RoleUpdateParams) (Role, error) {
	return r.post(ctx, fmt.Sprintf("/Services/%s/Roles/%s", sid, roleSID), body.encode())
}

// DELETE /Services/{Service SID}/Roles/{Role SID}
// https://www.twilio.com/docs/chat/rest/roles#delete-a-role
func (r roleAPI) Delete(ctx context.Context, sid, roleSID string) error {
	_, err := r.client.Delete(ctx, fmt.Sprintf("/Services/%s/Roles/%s", sid, roleSID))
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
