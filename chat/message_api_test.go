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

func TestMessageRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/csid/Messages/msid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/message.json")
		}

		var (
			exp  = Message{}
			f, _ = os.Open("fixtures/message.json")
		)
		json.NewDecoder(f).Decode(&exp)

		message, err := messageAPI{client}.Read(context.TODO(), "sid", "csid", "msid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, message) {
			t.Errorf("response diff %v", cmp.Diff(exp, message))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return messageAPI{client}.Read(ctx, "sid", "csid", "msid")
		}
		APIMock(fn).TestGets((t))
	})
}

func TestMessageSend(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/sid/Channels/csid/Messages"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(gotBody, expBody) {
				t.Errorf("exp body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/message.json")
		}

		var (
			exp  Message
			f, _ = os.Open("fixtures/message.json")
		)
		json.NewDecoder(f).Decode(&exp)

		message, err := messageAPI{client}.Send(context.TODO(), "sid", "csid", MessageCreateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, message) {
			t.Errorf("response diff %v", cmp.Diff(exp, message))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (messageAPI{client}).Send(ctx, "sid", "csid", MessageCreateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestMessageUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.PostFunc = func(ctx context.Context, path string, body io.Reader) ([]byte, error) {
			var (
				gotBody, _ = ioutil.ReadAll(body)
				expBody    = []byte("")
			)

			if exp := "/Services/sid/Channels/csid/Messages/msid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			if !bytes.Equal(expBody, gotBody) {
				t.Errorf("exp req body %s, got %s", expBody, gotBody)
			}
			return ioutil.ReadFile("fixtures/message.json")
		}
		var (
			exp  Message
			f, _ = os.Open("fixtures/message.json")
		)
		json.NewDecoder(f).Decode(&exp)

		role, err := (messageAPI{client}).Update(context.TODO(), "sid", "csid", "msid", MessageUpdateParams{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, role) {
			t.Errorf("response diff %v", cmp.Diff(exp, role))
		}
	})
	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (messageAPI{client}).Update(ctx, "sid", "csid", "msid", MessageUpdateParams{})
		}
		APIMock(fn).TestPosts((t))
	})
}

func TestUpdateDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.DeleteFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Channels/csid/Messages/msid"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return nil, nil
		}

		if err := (messageAPI{client}).Delete(context.TODO(), "sid", "csid", "msid"); err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !client.DeleteInvoked {
			t.Errorf("exp delete to have been invoked")
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			err := (messageAPI{client}).Delete(ctx, "sid", "csid", "msid")
			return nil, err
		}
		APIMock(fn).TestDeletes((t))
	})
}
