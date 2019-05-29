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
		// List(context.Context) ([]Service, error)
		Read(context.Context, string) (Service, error)
		Create(context.Context, ServiceUpdateParams) (Service, error)
		Update(context.Context, string, ServiceUpdateParams) (Service, error)
		Delete(context.Context, string) error
	}

	serviceAPI struct {
		client twilio.HTTPClient
	}
)

// TODO: alex
//func (api serviceAPI) List(ctx context.Context) ([]Service, error) {
//	var services []Service
//	data, err := api.client.Get(ctx, "/Services")
//	if err != nil {
//		return services, err
//	}
//	err = json.Unmarshal(data, &services)
//	return services, err
//}

func (api serviceAPI) Read(ctx context.Context, SID string) (Service, error) {
	var service Service
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s", SID))
	if err != nil {
		return service, err
	}
	err = json.Unmarshal(data, &service)
	return service, err
}

func (api serviceAPI) Create(ctx context.Context, service ServiceUpdateParams) (Service, error) {
	return api.post(ctx, "/Services", service)
}

func (api serviceAPI) Update(ctx context.Context, SID string, service ServiceUpdateParams) (Service, error) {
	return api.post(ctx, fmt.Sprintf("/Services/%s", SID), service)
}

func (api serviceAPI) Delete(ctx context.Context, SID string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s", SID))
	return err
}

func (api serviceAPI) post(ctx context.Context, path string, updateParams ServiceUpdateParams) (Service, error) {
	var updatedService Service
	body, _ := json.Marshal(updateParams) // err ignored as servceUpdateParams uses allowed types
	data, err := api.client.Post(ctx, path, bytes.NewReader(body))
	if err != nil {
		return updatedService, err
	}
	err = json.Unmarshal(data, &updatedService)
	return updatedService, err
}
