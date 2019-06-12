package chat

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBindingRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Bindings/bsid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/binding.json")
		}

		var (
			exp  = Binding{}
			f, _ = os.Open("fixtures/binding.json")
		)
		json.NewDecoder(f).Decode(&exp)

		binding, err := (bindingAPI{client}).Read(context.TODO(), "sid", "bsid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, binding) {
			t.Errorf("response diff %v", cmp.Diff(exp, binding))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (bindingAPI{client}).Read(ctx, "sid", "bsid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestBindingDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Bindings/bsid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		err := (bindingAPI{client}).Delete(context.TODO(), "sid", "bsid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !client.DeleteInvoked {
			t.Errorf(("exp httpclient.Delete to have been invoked"))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (bindingAPI{client}).Delete(ctx, "sid", "bsid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
