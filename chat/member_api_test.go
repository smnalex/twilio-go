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

func TestMemberRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/csid/Members/msid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/member.json")
		}

		var (
			exp  = Member{}
			f, _ = os.Open("fixtures/member.json")
		)
		json.NewDecoder(f).Decode(&exp)

		member, err := (memberAPI{client}).Read(context.TODO(), "sid", "csid", "msid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, member) {
			t.Errorf("response diff %v", cmp.Diff(exp, member))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (memberAPI{client}).Read(ctx, "sid", "csid", "msid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestMemberCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("Identity=")
			)

			if exp := "/Services/sid/Channels/csid/Members"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/member.json")
		}

		var (
			exp  Member
			f, _ = os.Open("fixtures/member.json")
		)
		json.NewDecoder(f).Decode(&exp)

		member, err := (memberAPI{client}).Add(context.TODO(), "sid", "csid", MemberCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, member) {
			t.Errorf("response diff %v", cmp.Diff(exp, member))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (memberAPI{client}).Add(ctx, "sid", "csid", MemberCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestMemberDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/csid/Members/msid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (memberAPI{client}).Delete(context.TODO(), "sid", "csid", "msid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}

		if !client.DeleteInvoked {
			t.Errorf(("exp channel.Delete to have been invoked"))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (memberAPI{client}).Delete(ctx, "sid", "csid", "msid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
