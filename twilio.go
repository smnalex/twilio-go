package twilio

import "net/http"

// Context store for credentials, configuration and the http client.
type Context struct {
	AccountSID string
	AuthToken  string
	Region     string
	HTTPClient RequestHandler
}

// NewContext store the various credentials into a `twilio.Context` instance and
// sets `http.DefaultClient` as HTTPClient
func NewContext(accountSID, authToken, region string) Context {
	return NewContextWithHTTP(accountSID, authToken, region, http.DefaultClient)
}

// NewContextWithHTTP sames as `NewContext` but requires a `twilio.RequestHandler`
func NewContextWithHTTP(accountSID, authToken, region string, httpClient RequestHandler) Context {
	return Context{
		AccountSID: accountSID,
		AuthToken:  authToken,
		Region:     region,
		HTTPClient: httpClient,
	}
}
