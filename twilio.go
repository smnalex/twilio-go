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
	AuthToken  string
	Region     string
	HTTPClient RequestHandler
}

// NewContext store the various credentials into a `twilio.Context` instance and
// sets `http.DefaultClient` as HTTPClient.
func NewContext() Context {
	return NewContextWithHTTP("", "", "", http.DefaultClient)
}

// NewContextWithHTTP sames as `NewContext` but requires a `twilio.RequestHandler`.
func NewContextWithHTTP(accountSID, authToken, region string, httpClient RequestHandler) Context {
	if accountSID == "" {
		accountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	}
	if authToken == "" {
		authToken = os.Getenv("TWILIO_SECRET_KEY")
	}
	if region == "" {
		region = os.Getenv("TWILIO_REGION")
	}
	return Context{
		AccountSID: accountSID,
		AuthToken:  authToken,
		Region:     region,
		HTTPClient: httpClient,
	}
}
