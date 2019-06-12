package chat

import "testing"

func TestMessageParamsOptionals(t *testing.T) {
	exp := []byte("")
	t.Run("CreateParams", optionalsFn(MessageCreateParams{}, exp))
	exp = []byte("")
	t.Run("UpdateParams", optionalsFn(MessageUpdateParams{}, exp))
}
