package chat

import (
	"testing"
)

func TestServiceParamsOptionals(t *testing.T) {
	exp := []byte("FriendlyName=")
	t.Run("CreateParams", optionalsFn(ServiceCreateParams{}, exp))
	exp = []byte("")
	t.Run("UpdateParams", optionalsFn(ServiceUpdateParams{}, exp))
	exp = []byte("Notifications.AddedToChannel.Enabled=true&Notifications.LogEnabled=true&Notifications.NewMessage.BadgeCountEnabled=true")
	t.Run("UpdateParams", optionalsFn(ServiceUpdateParams{
		Notifications: &Notifications{
			LogEnabled: true,
			AddedToChannel: &NotificationChannelProperty{
				Enabled: true,
			},
			NewMessage: &NotificationChannelProperty{
				BadgeCountEnabled: true,
			},
		},
	}, exp))
}
