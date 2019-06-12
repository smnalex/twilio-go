package twilio

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// HTTPClient provides http methods required for making requests to the Twilio API.
type HTTPClient interface {
	Get(context.Context, string) ([]byte, error)
	Post(context.Context, string, io.Reader) ([]byte, error)
	Delete(context.Context, string) ([]byte, error)
}

// RequestHandler defines the method required for an HTTP client to execute requests.
// An *http.Client satisfies this interface.
type RequestHandler interface {
	Do(*http.Request) (*http.Response, error)
}

type httpClient struct {
	url       *url.URL
	apiKey    string
	apiSecret string
	RequestHandler
}

// NewHTTPClient returns a new HTTPClient customised for making Twilio http requests.
func NewHTTPClient(apiKey, apiSecret, baseURL string, rh RequestHandler) (HTTPClient, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse url")
	}

	return &httpClient{
		url:            url,
		apiKey:         apiKey,
		apiSecret:      apiSecret,
		RequestHandler: rh,
	}, nil
}

func (client *httpClient) Get(ctx context.Context, path string) ([]byte, error) {
	return client.request(ctx, http.MethodGet, path, nil)
}

func (client *httpClient) Post(ctx context.Context, path string, body io.Reader) ([]byte, error) {
	return client.request(ctx, http.MethodPost, path, body)
}

func (client *httpClient) Delete(ctx context.Context, path string) ([]byte, error) {
	return client.request(ctx, http.MethodDelete, path, nil)
}

func (client *httpClient) request(ctx context.Context, method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, client.url.String()+path, body)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not create request")
	}

	{
		req.SetBasicAuth(client.apiKey, client.apiSecret)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req = req.WithContext(ctx)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "httpclient: could not get a response for %s", req.URL)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode >= http.StatusBadRequest {
		return nil, decodeErr(resp.Body)
	}

	return ioutil.ReadAll(resp.Body)
}

func decodeErr(b io.Reader) error {
	var err ErrTwilioResponse
	if err := json.NewDecoder(b).Decode(&err); err != nil {
		return err
	}
	return err
}
