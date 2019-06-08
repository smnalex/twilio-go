package chat

import "context"

type (
	// UserChannelRepository interface for interacting with the user channels chat api.
	UserChannelRepository interface {
		Read(ctx context.Context, serviceSid, userSid string)
	}
)
