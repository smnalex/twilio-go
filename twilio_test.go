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
		auth = "auth"
		region = "region"
	)
 setup := func() func() {
	os.Setenv("TWILIO_ACCOUNT_SID", acc)
	os.Setenv("TWILIO_SECRET_KEY", auth)
	os.Setenv("TWILIO_REGION", region)

	return func() {
		os.Unsetenv("TWILIO_ACCOUNT_SID")
		os.Unsetenv("TWILIO_SECRET_KEY")
		os.Unsetenv("TWILIO_REGION")
	}
}
cleanup := setup()
defer cleanup()
	c := NewContext()

	if c.AccountSID != acc {
		t.Errorf("exp accsid %v, got %v", acc, c.AccountSID)
	}
	if c.AuthToken != auth {
		t.Errorf("exp auth %v, got %v", auth, c.AuthToken)
	}
	if c.Region != region {
		t.Errorf("exp auth %v, got %v", auth, c.AuthToken)
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