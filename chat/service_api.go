package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/smnalex/twilio-go"
)

type serviceAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}
// https://www.twilio.com/docs/chat/rest/services#retrieve-a-service
func (api serviceAPI) Read(ctx context.Context, serviceSid string) (Service, error) {
	var service Service
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s", serviceSid))
	if err != nil {
		return service, err
	}
	err = json.Unmarshal(data, &service)
	return service, err
}

// POST /Services
// https://www.twilio.com/docs/chat/rest/services#create-a-service
func (api serviceAPI) Create(ctx context.Context, body ServiceCreateParams) (Service, error) {
	return api.post(ctx, "/Services", body.encode())
}

// POST /Services/{Service SID}
// https://www.twilio.com/docs/chat/rest/services#update-a-service
func (api serviceAPI) Update(ctx context.Context, serviceSid string, body ServiceUpdateParams) (Service, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s", serviceSid), body.encode())
}

// DELETE /Services/{Service SID}
// https://www.twilio.com/docs/chat/rest/services#delete-a-service
func (api serviceAPI) Delete(ctx context.Context, serviceSid string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s", serviceSid))
	return err
}

func (api serviceAPI) post(ctx context.Context, path string, body io.Reader) (Service, error) {
	var service Service
	data, err := api.client.Post(ctx, path, body)
	if err != nil {
		return service, err
	}
	err = json.Unmarshal(data, &service)
	return service, err
}
