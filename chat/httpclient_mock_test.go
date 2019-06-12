package chat

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/smnalex/twilio-go"
)

type APIMock func(context.Context, *HTTPClientMock) (interface{}, error)

func (triggerFn APIMock) TestGets(t *testing.T) {
	ctx := context.Background()
	t.Run("response parsing error", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			return []byte("invalid"), nil
		}

		if _, err := triggerFn(ctx, client); err == nil {
			t.Errorf("exp parsing err, got %v", err)
		}
	})
	t.Run("api response error", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			return nil, twilio.ErrTwilioResponse{}
		}

		exp := twilio.ErrTwilioResponse{}
		if _, err := triggerFn(ctx, client); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
	t.Run("api request ctx timeout", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			select {
			case <-time.After(time.Second * 1):
				break
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			return nil, nil
		}
		ctx, cancelFn := context.WithTimeout(ctx, 1*time.Microsecond)
		defer cancelFn()

		exp := context.DeadlineExceeded
		if _, err := triggerFn(ctx, client); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func (triggerFn APIMock) TestPosts(t *testing.T) {
	ctx := context.Background()
	t.Run("response parsing error", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			return []byte("invalid"), nil
		}

		if _, err := triggerFn(ctx, client); err == nil {
			t.Errorf("exp parsing err, got %v", err)
		}
	})
	t.Run("api response error", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			return nil, twilio.ErrTwilioResponse{}
		}

		exp := twilio.ErrTwilioResponse{}
		if _, err := triggerFn(ctx, client); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
	t.Run("api request ctx timeout", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			select {
			case <-time.After(time.Second * 1):
				break
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			return nil, nil
		}

		ctx, cancelFn := context.WithTimeout(ctx, 1*time.Microsecond)
		defer cancelFn()

		exp := context.DeadlineExceeded
		if _, err := triggerFn(ctx, client); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func (triggerFn APIMock) TestDeletes(t *testing.T) {
	ctx := context.Background()
	t.Run("api response error", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			return nil, twilio.ErrTwilioResponse{}
		}

		exp := twilio.ErrTwilioResponse{}
		if _, err := triggerFn(ctx, client); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
	t.Run("api request ctx timeout", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			select {
			case <-time.After(time.Second * 1):
				break
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			return nil, nil
		}

		ctx, cancel := context.WithTimeout(ctx, 1*time.Microsecond)
		defer cancel()

		exp := context.DeadlineExceeded
		if _, err := triggerFn(ctx, client); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

type HTTPClientMock struct {
	GetFunc       func(context.Context, string) ([]byte, error)
	PostFunc      func(context.Context, string, io.Reader) ([]byte, error)
	DeleteInvoked bool
	DeleteFunc    func(context.Context, string) ([]byte, error)
}

func (m *HTTPClientMock) Get(ctx context.Context, path string) ([]byte, error) {
	return m.GetFunc(ctx, path)
}

func (m *HTTPClientMock) Post(ctx context.Context, path string, body io.Reader) ([]byte, error) {
	return m.PostFunc(ctx, path, body)
}

func (m *HTTPClientMock) Delete(ctx context.Context, path string) ([]byte, error) {
	m.DeleteInvoked = true
	return m.DeleteFunc(ctx, path)
}
