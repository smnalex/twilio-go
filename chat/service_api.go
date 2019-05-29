package chat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/smnalex/twilio-go"
)

type (
	// ServiceRepository defines the interface for interacting with chat service api.
	ServiceRepository interface {
		Read(context.Context, string) (Service, error)
		Create(context.Context, string) (Service, error)
		Update(context.Context, string, ServiceUpdateParams) (Service, error)
		Delete(context.Context, string) error
	}

	serviceAPI struct {
		client twilio.HTTPClient
	}
)

// GET /Services/{Service SID}
// https://www.twilio.com/docs/chat/rest/services#retrieve-a-service
func (api serviceAPI) Read(ctx context.Context, SID string) (Service, error) {
	var service Service
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s", SID))
	if err != nil {
		return service, err
	}
	err = json.Unmarshal(data, &service)
	return service, err
}

// POST /Services
// https://www.twilio.com/docs/chat/rest/services#create-a-service
func (api serviceAPI) Create(ctx context.Context, friendlyName string) (Service, error) {
	return api.post(ctx, "/Services", ServiceUpdateParams{FriendlyName: friendlyName})
}

// POST /Services/{Service SID}
// https://www.twilio.com/docs/chat/rest/services#update-a-service
func (api serviceAPI) Update(ctx context.Context, SID string, service ServiceUpdateParams) (Service, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s", SID), service)
}

// DELETE /Services/{Service SID}
// https://www.twilio.com/docs/chat/rest/services#delete-a-service
func (api serviceAPI) Delete(ctx context.Context, SID string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s", SID))
	return err
}

func (api serviceAPI) post(ctx context.Context, path string, updateParams ServiceUpdateParams) (Service, error) {
	var updatedService Service
	body, _ := json.Marshal(updateParams) // err ignored as `ServiceUpdateParams` uses allowed types
	data, err := api.client.Post(ctx, path, bytes.NewReader(body))
	if err != nil {
		return updatedService, err
	}
	err = json.Unmarshal(data, &updatedService)
	return updatedService, err
}
