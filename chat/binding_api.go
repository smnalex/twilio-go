package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smnalex/twilio-go"
)

type bindingAPI struct {
	client twilio.HTTPClient
}

// GET /Services/{Service SID}/Bindings/{Binding SID}
// https://www.twilio.com/docs/chat/rest/bindings-resource#read-a-binding
func (api bindingAPI) Read(ctx context.Context, serviceSid, bindingSid string) (Binding, error) {
	var bind Binding
	data, err := api.client.Get(ctx, fmt.Sprintf("/Services/%s/Bindings/%s", serviceSid, bindingSid))
	if err != nil {
		return bind, err
	}
	err = json.Unmarshal(data, &bind)
	return bind, err
}

func (api bindingAPI) Delete(ctx context.Context, serviceSid, bindingSid string) error {
	_, err := api.client.Delete(ctx, fmt.Sprintf("/Services/%s/Bindings/%s", serviceSid, bindingSid))
	return err
}
