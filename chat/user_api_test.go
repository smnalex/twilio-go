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

func TestUserRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Users/identity"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/user.json")
		}

		var (
			exp  = User{}
			f, _ = os.Open("fixtures/user.json")
		)
		json.NewDecoder(f).Decode(&exp)

		user, err := (userAPI{client}).Read(context.TODO(), "sid", "identity")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, user) {
			t.Errorf("response diff %v", cmp.Diff(exp, user))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (userAPI{client}).Read(ctx, "sid", "identity")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestUserCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("Identity=")
			)

			if exp := "/Services/sid/Users"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/user.json")
		}

		var (
			exp  User
			f, _ = os.Open("fixtures/user.json")
		)
		json.NewDecoder(f).Decode(&exp)

		user, err := (userAPI{client}).Create(context.TODO(), "sid", UserCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, user) {
			t.Errorf("response diff %v", cmp.Diff(exp, user))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (userAPI{client}).Create(ctx, "sid", UserCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestUserUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/sid/Users/usid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(expBody, gotBody) {
				t.Errorf("exp req body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/user.json")
		}

		var (
			exp  User
			f, _ = os.Open("fixtures/user.json")
		)
		json.NewDecoder(f).Decode(&exp)

		user, err := (userAPI{client}).Update(context.TODO(), "sid", "usid", UserUpdateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, user) {
			t.Errorf("response diff %v", cmp.Diff(exp, user))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (userAPI{client}).Update(ctx, "sid", "usid", UserUpdateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestUserDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Users/usid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (userAPI{client}).Delete(context.TODO(), "sid", "usid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (userAPI{client}).Delete(ctx, "sid", "usid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
