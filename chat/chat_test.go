package chat

import (
	"os"
	"testing"

	"github.com/smnalex/twilio-go"
)

func TestNew(t *testing.T) {
	t.Run("unsuccessful invalid env url", func(t *testing.T) {
		os.Setenv("TWILIO_CHAT_HOST", "%2")
		if _, err := New(twilio.Context{}); err == nil {
			t.Errorf("exp parsing err, got none")
		}
		os.Unsetenv("TWILIO_CHAT_HOST")
	})

	t.Run("chat services", func(t *testing.T) {
		_, err := New(twilio.Context{})
		if err != nil {
			t.Errorf("exp no err, got %v", err)
		}
	})
}

func TestChatEndpoint(t *testing.T) {
	exp := "https://chat.twilio.com/v2"

	t.Run("default url", func(*testing.T) {
		if got := chatEndpointForRegion(""); got != exp {
			t.Errorf("exp url %s, got %s", exp, got)
		}
	})

	t.Run("default url with region", func(*testing.T) {
		exp := "https://chat.uk.twilio.com/v2"
		if got := chatEndpointForRegion("uk"); got != exp {
			t.Errorf("exp url %s, got %s", exp, got)
		}
	})

	t.Run("env url", func(*testing.T) {
		os.Setenv("TWILIO_CHAT_HOST", exp)
		if got := chatEndpointForRegion(""); got != exp {
			t.Errorf("exp url %s, got %s", exp, got)
		}
		if got := chatEndpointForRegion("uk"); got != exp {
			t.Errorf("exp url %s, got %s", exp, got)
		}
		os.Unsetenv("TWILIO_CHAT_HOST")
	})
}
