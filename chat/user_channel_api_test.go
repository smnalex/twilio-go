package chat

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUserChannelList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := &HTTPClientMock{}
		client.GetFunc = func(ctx context.Context, path string) ([]byte, error) {
			if exp := "/Services/sid/Users/usid/Channels"; exp != path {
				t.Errorf("exp path %s, got %s", exp, path)
			}
			return ioutil.ReadFile("fixtures/user_channels.json")
		}

		var (
			exp  = UserChannelList{}
			f, _ = os.Open("fixtures/user_channels.json")
		)
		json.NewDecoder(f).Decode(&exp)

		userChannel, err := (userChannelAPI{client}).List(context.TODO(), "sid", "usid")
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
		if !cmp.Equal(exp, userChannel) {
			t.Errorf("response diff %v", cmp.Diff(exp, userChannel))
		}
	})

	t.Run("errors", func(t *testing.T) {
		fn := func(ctx context.Context, client *HTTPClientMock) (interface{}, error) {
			return (userChannelAPI{client}).List(ctx, "sid", "usid")
		}
		APIMock(fn).TestGets((t))
	})
}
