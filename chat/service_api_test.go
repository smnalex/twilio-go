package chat

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestServiceRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/service.json")
		}

		var (
			exp  = Service{}
			f, _ = os.Open("fixtures/service.json")
		)
		json.NewDecoder(f).Decode(&exp)

		service, err := (serviceAPI{client}).Read(context.TODO(), "sid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, service) {
			t.Errorf("response diff %v", cmp.Diff(exp, service))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (serviceAPI{client}).Read(ctx, "sid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestServiceCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
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

		service, err := (serviceAPI{client}).Create(context.TODO(), ServiceCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(service, exp) {
			t.Errorf("response diff %v", cmp.Diff(exp, service))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (serviceAPI{client}).Create(ctx, ServiceCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/sid"; exp != path {
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

		service, err := (serviceAPI{client}).Update(context.TODO(), "sid", ServiceUpdateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, service) {
			t.Errorf("response diff %v", cmp.Diff(exp, service))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (serviceAPI{client}).Update(ctx, "sid", ServiceUpdateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (serviceAPI{client}).Delete(context.TODO(), "sid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (serviceAPI{client}).Delete(ctx, "sid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
