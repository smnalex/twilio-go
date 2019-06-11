package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type credentialAPI struct {
	client twilio.HTTPClient
}

// GET /Credentials/{Credential SID}
// https://www.twilio.com/docs/chat/rest/credentials#retrieve-a-credential
func (api credentialAPI) Read(ctx context.Context, credentialSID string) (Credential, error) {
	var crd Credential
	data, err := api.client.Get(ctx, fmt.Sprintf("/Credentials/%s", credentialSID))
	if err != nil {
		return crd, err
	}
	err = json.Unmarshal(data, &crd)
	return crd, err
}

// POST /Credentials
// https://www.twilio.com/docs/chat/rest/credentials#create-a-credential
func (api credentialAPI) Create(ctx context.Context, body CredentialCreateParams) (Credential, error) {
	return api.post(ctx, "/Credentials", body.encode())
}

// POST /Credentials/{Credential SID}
// https://www.twilio.com/docs/chat/rest/credentials#update-a-credential
func (api credentialAPI) Update(ctx context.Context, credentialSID string, body CredentialUpdateParams) (Credential, error) {
	return api.post(ctx, fmt.Sprintf("/Credentials/%s", credentialSID), body.encode())
}

// DELETE /Credentials/{Credential SID}
// https://www.twilio.com/docs/chat/rest/credentials#delete-a-credential
func (api credentialAPI) Delete(ctx context.Context, credentialSID string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Credentials/%s", credentialSID))
	return err
}

func (api credentialAPI) post(ctx context.Context, path string, body io.Reader) (Credential, error) {
	var crd Credential
	data, err := api.client.Post(ctx, path, body)
	if err != nil {
		return crd, err
	}
	err = json.Unmarshal(data, &crd)
	return crd, err
}
