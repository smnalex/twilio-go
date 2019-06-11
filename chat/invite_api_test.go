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

func TestInviteRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/csid/Invites/isid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/invite.json")
		}

		var (
			exp  = Invite{}
			f, _ = os.Open("fixtures/invite.json")
		)
		json.NewDecoder(f).Decode(&exp)

		channel, err := (inviteAPI{client}).Read(context.TODO(), "sid", "csid", "isid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, channel) {
			t.Errorf("response diff %v", cmp.Diff(exp, channel))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (inviteAPI{client}).Read(ctx, "sid", "csid", "isid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestInviteCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("Identity=")
			)

			if exp := "/Services/sid/Channels/csid/Invites"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/invite.json")
		}

		var (
			exp  Invite
			f, _ = os.Open("fixtures/invite.json")
		)
		json.NewDecoder(f).Decode(&exp)

		invite, err := (inviteAPI{client}).Create(context.TODO(), "sid", "csid", InviteCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, invite) {
			t.Errorf("response diff %v", cmp.Diff(exp, invite))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (inviteAPI{client}).Create(ctx, "sid", "csid", InviteCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestInviteDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/csid/Invites/isid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (inviteAPI{client}).Delete(context.TODO(), "sid", "csid", "isid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}

		if !client.DeleteInvoked {
			t.Errorf(("exp channel.Delete to have been invoked"))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (inviteAPI{client}).Delete(ctx, "sid", "csid", "isid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
