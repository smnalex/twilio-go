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

func TestChannelRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/identity"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/channel.json")
		}

		var (
			exp  = Channel{}
			f, _ = os.Open("fixtures/channel.json")
		)
		json.NewDecoder(f).Decode(&exp)

		channel, err := (channelAPI{client}).Read(context.TODO(), "sid", "identity")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, channel) {
			t.Errorf("response diff %v", cmp.Diff(exp, channel))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (channelAPI{client}).Read(ctx, "sid", "identity")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestChannelCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/sid/Channels"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/channel.json")
		}

		var (
			exp  Channel
			f, _ = os.Open("fixtures/channel.json")
		)
		json.NewDecoder(f).Decode(&exp)

		channel, err := (channelAPI{client}).Create(context.TODO(), "sid", ChannelCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, channel) {
			t.Errorf("response diff %v", cmp.Diff(exp, channel))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (channelAPI{client}).Create(ctx, "sid", ChannelCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestChannelUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/sid/Channels/identity"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(expBody, gotBody) {
				t.Errorf("exp req body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/channel.json")
		}

		var (
			exp  Channel
			f, _ = os.Open("fixtures/channel.json")
		)
		json.NewDecoder(f).Decode(&exp)

		channel, err := (channelAPI{client}).Update(context.TODO(), "sid", "identity", ChannelUpdateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, channel) {
			t.Errorf("response diff %v", cmp.Diff(exp, channel))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (channelAPI{client}).Update(ctx, "sid", "identity", ChannelUpdateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestChannelDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/identity"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (channelAPI{client}).Delete(context.TODO(), "sid", "identity"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}

		if !client.DeleteInvoked {
			t.Errorf(("exp channel.Delete to have been invoked"))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (channelAPI{client}).Delete(ctx, "sid", "identity")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
