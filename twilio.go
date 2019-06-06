package twilio

import (
	"fmt"
	"net/http"
	"os"
)

// ErrTwilioResponse returned when response codes are greater than 400.
type ErrTwilioResponse struct {
	Code    int
	Status  int
	Message string
}

func (e ErrTwilioResponse) Error() string {
	return fmt.Sprintf("%d: %d, %s", e.Status, e.Code, e.Message)
}

// Context store for credentials, configuration and the http client.
type Context struct {
	AccountSID string

	// APIKey contains a secret used to sign Access Tokens.
	APIKey    string
	APISecret string

	Region string

	HTTPClient RequestHandler
}

// NewContext returns a new Context with a http.DefaultClient and various informations
// loaded from envs.
func NewContext() Context {
	return NewContextWithHTTP("", "", "", "", http.DefaultClient)
}

// NewContextWithHTTP sames as `NewContext` but requires a `twilio.RequestHandler`.
func NewContextWithHTTP(accountSID, apiKey, apiSecret, region string, httpClient RequestHandler) Context {
	if accountSID == "" {
		accountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	}
	if apiKey == "" {
		apiKey = os.Getenv("TWILIO_API_KEY")
	}
	if apiSecret == "" {
		apiSecret = os.Getenv("TWILIO_API_SECRET_KEY")
	}
	if region == "" {
		region = os.Getenv("TWILIO_API_REGION")
	}
	return Context{
		AccountSID: accountSID,
		APIKey:     apiKey,
		APISecret:  apiSecret,
		Region:     region,
		HTTPClient: httpClient,
	}
}
