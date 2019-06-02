package chat

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/smnalex/twilio-go"
)

func TestServiceRead(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		services   serviceAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			services = serviceAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/SID"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/service.json")
		}

		var (
			exp  = Service{}
			f, _ = os.Open("fixtures/service.json")
		)
		json.NewDecoder(f).Decode(&exp)

		service, err := services.Read(ctx, "SID")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, service) {
			t.Errorf("response diff %v", cmp.Diff(exp, service))
		}
	})

	t.Run("response parsing body error", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			return []byte("invalid"), nil
		}

		if _, err := services.Read(ctx, "SID"); err == nil {
			t.Errorf("exp parsing err, got %v", err)
		}
	})

	t.Run("api response error", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			return nil, twilio.ErrTwilioResponse{}
		}

		exp := twilio.ErrTwilioResponse{}
		if _, err := services.Read(ctx, "SID"); exp != err {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})

	t.Run("api request ctx timeout", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
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
		if _, err := services.Read(ctx, "SID"); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func TestServiceCreate(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		services   serviceAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			services = serviceAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("FriendlyName=")
			)

			if exp := "/Services"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/service.json")
		}

		var (
			exp  Service
			f, _ = os.Open("fixtures/service.json")
		)
		json.NewDecoder(f).Decode(&exp)

		service, err := services.Create(ctx, ServiceCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(service, exp) {
			t.Errorf("response diff %v", cmp.Diff(exp, service))
		}
	})

	t.Run("response parsing body error", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			return []byte("invalid"), nil
		}

		if _, err := services.Create(ctx, ServiceCreateParams{}); err == nil {
			t.Errorf("exp parsing err, got %v", err)
		}
	})

	t.Run("api request ctx timeout", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
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
		if _, err := services.Create(ctx, ServiceCreateParams{}); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func TestServiceUpdate(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		services   serviceAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			services = serviceAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/SID"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(expBody, gotBody) {
				t.Errorf("exp req body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/service.json")
		}

		var (
			exp  Service
			f, _ = os.Open("fixtures/service.json")
		)
		json.NewDecoder(f).Decode(&exp)

		service, err := services.Update(ctx, "SID", ServiceUpdateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, service) {
			t.Errorf("response diff %v", cmp.Diff(exp, service))
		}
	})

	t.Run("api request ctx timeout", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
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
		if _, err := services.Update(ctx, "SID", ServiceUpdateParams{}); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func TestServiceDelete(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		services   serviceAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			services = serviceAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/SID"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := services.Delete(ctx, "SID"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
	})

	t.Run("api request ctx timeout", func(t *testing.T) {
		setup()
		mockClient.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			select {
			case <-time.After(1 * time.Second):
				break
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			return nil, nil
		}
		ctx, cancel := context.WithTimeout(ctx, 5*time.Microsecond)
		defer cancel()

		exp := context.DeadlineExceeded
		if err := services.Delete(ctx, "SID"); err != exp {
			t.Errorf("exp err %v err, got %v", exp, err)
		}
	})
}
