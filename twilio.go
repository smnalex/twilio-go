package twilio

import "net/http"

// TwilioContext store for credentials, configuration and the http client.
type TwilioContext struct {
	AccountSID string
	AuthToken  string
	Region     string
	HTTPClient RequestHandler
}

// NewContext store the various credentials into a twilio instance which comes
// by default with a `http.DefaultClient`.
func NewContext(accountSID, authToken, region string) TwilioContext {
	return NewContextWithHTTP(accountSID, authToken, region, http.DefaultClient)
}

// NewContextWithHTTP sames as `NewContext` but requires a `RequestHandler`
func NewContextWithHTTP(accountSID, authToken, region string, httpClient RequestHandler) TwilioContext {
	return TwilioContext{
		AccountSID: accountSID,
		AuthToken:  authToken,
		Region:     region,
		HTTPClient: httpClient,
	}
}
