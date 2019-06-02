package chat

import (
	"testing"
)

func TestServiceParamsOptionals(t *testing.T) {
	t.Run("CreateParams", func(t *testing.T) {
		exp := []byte("FriendlyName=")
		t.Run("CreateParams", optionalsFn(ServiceCreateParams{}, exp))
		exp = []byte("")
		t.Run("UpdateParams", optionalsFn(ServiceUpdateParams{}, exp))
	})
}
