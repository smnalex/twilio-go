package twilio

import (
	"os"
	"testing"
	"net/http"
	"fmt"
)

func TestNewContext(t *testing.T) {
	var (
		acc = "acc"
		apiKey = "auth"
		apiSecret = "secret"
		region = "region"

		setup = func() func() {
			os.Setenv("TWILIO_ACCOUNT_SID", acc)
			os.Setenv("TWILIO_API_KEY", apiKey)
			os.Setenv("TWILIO_API_SECRET_KEY", apiSecret)
			os.Setenv("TWILIO_API_REGION", region)

			return func() {
				os.Unsetenv("TWILIO_ACCOUNT_SID")
				os.Unsetenv("TWILIO_API_KEY")
				os.Unsetenv("TWILIO_API_SECRET_KEY")
				os.Unsetenv("TWILIO_REGION")
			}
		}
		cleanup = setup()
	)

	defer cleanup()
	c := NewContext()

	if c.AccountSID != acc {
		t.Errorf("exp accsid %v, got %v", acc, c.AccountSID)
	}
	if c.APIKey != apiKey {
		t.Errorf("exp api key %v, got %v", apiKey, c.APIKey)
	}
	if c.APISecret != apiSecret {
		t.Errorf("exp api secret %s, got %s", apiSecret, c.APISecret)
	}
	if c.Region != region {
		t.Errorf("exp auth %s, got %s", auth, c.Region)
	}
	if c.HTTPClient != http.DefaultClient {
		t.Errorf("exp *http.Client, got %T", c.HTTPClient)
	}
}

func TestErrTwilioResponse(t *testing.T) {
	err := ErrTwilioResponse{1, 2, "msg"}

	exp := fmt.Sprintf("%d: %d, %s", err.Status, err.Code, err.Message)
	if exp != err.Error() {
		t.Errorf("exp err msg %s, got %s", exp, err.Error())
	}
}