package chat

// UserChannelResource handles interactions with User Channels Programmable Chat REST API.
type UserChannelResource struct {
	userChannelAPI
}

// UserChannelList resource of Programmable Chat is a list only resource which will return
// the list of Channels the User is a Member of.
type UserChannelList struct {
	Channels []UserChannel `json:"channels"`
	Meta     Meta          `json:"meta"`
}

// UserChannel represents a channel the User is a Member of.
type UserChannel struct {
	AccountSid               string `json:"account_sid"`
	ServiceSid               string `json:"service_sid"`
	ChannelSid               string `json:"channel_sid"`
	UserSid                  string `json:"user_sid"`
	MemberSid                string `json:"member_sid"`
	Status                   string `json:"status"`
	LastConsumedMessageIndex int    `json:"last_consumed_message_index"`
	UnreadMessagesCount      int    `json:"unread_messages_count"`
	NotificationLevel        string `json:"notification_level"`
	URL                      string `json:"url"`
	Links                    struct {
		Channel string `json:"channel"`
		Member  string `json:"member"`
	} `json:"links"`
}
