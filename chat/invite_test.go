package chat

import "testing"

func TestInviteParamsOptionals(t *testing.T) {
	exp := []byte("Identity=")
	t.Run("CreateParams", optionalsFn(InviteCreateParams{}, exp))
}
