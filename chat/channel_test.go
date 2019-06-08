package chat

import "testing"

func TestChannelParamsOptionals(t *testing.T) {
	exp := []byte("")
	t.Run("CreateParams", optionalsFn(ChannelCreateParams{}, exp))
	exp = []byte("")
	t.Run("UpdateParams", optionalsFn(ChannelUpdateParams{}, exp))
}
