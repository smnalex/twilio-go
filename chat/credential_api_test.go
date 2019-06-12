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

func TestCredentialRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Credentials/csid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/role.json")
		}

		var (
			exp  = Credential{}
			f, _ = os.Open("fixtures/role.json")
		)
		json.NewDecoder(f).Decode(&exp)

		crd, err := credentialAPI{client}.Read(context.TODO(), "csid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, crd) {
			t.Errorf("response diff %v", cmp.Diff(exp, crd))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return credentialAPI{client}.Read(ctx, "csid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestCredentialCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("Type=")
			)

			if exp := "/Credentials"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/credential.json")
		}

		var (
			exp  Credential
			f, _ = os.Open("fixtures/credential.json")
		)
		json.NewDecoder(f).Decode(&exp)

		crd, err := credentialAPI{client}.Create(context.TODO(), CredentialCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, crd) {
			t.Errorf("response diff %v", cmp.Diff(exp, crd))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (credentialAPI{client}).Create(ctx, CredentialCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestCredentialUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Credentials/csid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(expBody, gotBody) {
				t.Errorf("exp req body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/credential.json")
		}
		var (
			exp  Credential
			f, _ = os.Open("fixtures/credential.json")
		)
		json.NewDecoder(f).Decode(&exp)

		crd, err := (credentialAPI{client}).Update(context.TODO(), "csid", CredentialUpdateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, crd) {
			t.Errorf("response diff %v", cmp.Diff(exp, crd))
		}
	})
	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (credentialAPI{client}).Update(ctx, "csid", CredentialUpdateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestCredentialDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Credentials/csid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (credentialAPI{client}).Delete(context.TODO(), "csid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !client.DeleteInvoked {
			t.Errorf("exp delete to have been invoked")
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (credentialAPI{client}).Delete(ctx, "csid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
