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

func TestRoleRead(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		roles      roleAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			roles = roleAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/SID/Roles/ROLESID"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/role.json")
		}

		var (
			exp  = Role{}
			f, _ = os.Open("fixtures/role.json")
		)
		json.NewDecoder(f).Decode(&exp)

		role, err := roles.Read(ctx, "SID", "ROLESID")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
		}
	})

	t.Run("response parsing body error", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			return []byte("invalid"), nil
		}

		if _, err := roles.Read(ctx, "SID", "ROLESID"); err == nil {
			t.Errorf("exp parsing err, got %v", err)
		}
	})

	t.Run("api response error", func(t *testing.T) {
		setup()
		mockClient.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			return nil, twilio.ErrTwilioResponse{}
		}

		exp := twilio.ErrTwilioResponse{}
		if _, err := roles.Read(ctx, "SID", "ROLESID"); err != exp {
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
		if _, err := roles.Read(ctx, "SID", "ROLESID"); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func TestRoleCreate(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		roles      roleAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			roles = roleAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("FriendlyName=&Type=")
			)

			if exp := "/Services/SID/Roles"; exp != path {
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

		role, err := roles.Create(ctx, "SID", RoleCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
		}
	})

	t.Run("response parsing body error", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			return []byte("invalid"), nil
		}

		if _, err := roles.Create(ctx, "SID", RoleCreateParams{}); err == nil {
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
		if _, err := roles.Create(ctx, "SID", RoleCreateParams{}); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func TestRoleUpdate(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		roles      roleAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			roles = roleAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("Permission=A&Permission=B")
			)

			if exp := "/Services/SID/Roles/ROLESID"; exp != path {
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

		role, err := roles.Update(ctx, "SID", "ROLESID", RoleUpdateParams{[]string{"A", "B"}})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
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
		if _, err := roles.Update(ctx, "SID", "ROLESID", RoleUpdateParams{}); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}

func TestRoleDelete(t *testing.T) {
	var (
		mockClient *mockHTTPClient
		roles      roleAPI

		ctx   = context.Background()
		setup = func() {
			mockClient = &mockHTTPClient{}
			roles = roleAPI{mockClient}
		}
	)

	t.Run("success", func(t *testing.T) {
		setup()
		mockClient.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/SID/Roles/ROLESID"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}
		if err := roles.Delete(ctx, "SID", "ROLESID"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
	})

	t.Run("api request ctx timeout", func(t *testing.T) {
		setup()
		mockClient.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
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
		if err := roles.Delete(ctx, "SID", "ROLESID"); err != exp {
			t.Errorf("exp err %v, got %v", exp, err)
		}
	})
}
