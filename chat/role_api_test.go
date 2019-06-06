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

func TestRoleRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Roles/rolesid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/role.json")
		}

		var (
			exp  = Role{}
			f, _ = os.Open("fixtures/role.json")
		)
		json.NewDecoder(f).Decode(&exp)

		role, err := roleAPI{client}.Read(context.TODO(), "sid", "rolesid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return roleAPI{client}.Read(ctx, "sid", "roleSid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestRoleCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("FriendlyName=&Type=")
			)

			if exp := "/Services/sid/Roles"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/role.json")
		}

		var (
			exp  Role
			f, _ = os.Open("fixtures/role.json")
		)
		json.NewDecoder(f).Decode(&exp)

		role, err := roleAPI{client}.Create(context.TODO(), "sid", RoleCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (roleAPI{client}).Create(ctx, "sid", RoleCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestRoleUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("Permission=A&Permission=B")
			)

			if exp := "/Services/sid/Roles/rolesid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(expBody, gotBody) {
				t.Errorf("exp req body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/role.json")
		}
		var (
			exp  Role
			f, _ = os.Open("fixtures/role.json")
		)
		json.NewDecoder(f).Decode(&exp)

		role, err := (roleAPI{client}).Update(context.TODO(), "sid", "rolesid", RoleUpdateParams{[]string{"A", "B"}})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
		}
	})
	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (roleAPI{client}).Update(ctx, "sid", "roleSid", RoleUpdateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestRoleDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Roles/rolesid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (roleAPI{client}).Delete(context.TODO(), "sid", "rolesid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !client.DeleteInvoked {
			t.Errorf("exp delete to have been invoked")
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (roleAPI{client}).Delete(ctx, "sid", "rolesid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
