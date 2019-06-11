package chat

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/go-querystring/query"
)

// ChannelResource handles interactions with Channels Programmable Chat REST API.
type ChannelResource struct {
	channelAPI
}

// Channel resource of Programmable Chat represents a "chat room".
type Channel struct {
	Sid          string          `json:"sid"`
	AccountSid   string          `json:"account_sid"`
	ServiceSid   string          `json:"service_sid"`
	FriendlyName string          `json:"friendly_name"`
	UniqueName   string          `json:"unique_name"`
	Attributes   json.RawMessage `json:"attributes"`
	Type         string          `json:"type"`

	// DateCreated RFC 2822 format.
	DateCreated string `json:"date_created"`

	// DateUpdated RFC 2822 format.
	DateUpdated   string `json:"date_updated"`
	CreatedBy     string `json:"created_by"`
	MembersCount  int    `json:"members_count"`
	MessagesCount int    `json:"messages_count"`
	URL           string `json:"url"`
	Links         struct {
		Members     string      `json:"members"`
		Messages    string      `json:"messages"`
		Invites     string      `json:"invites"`
		Webhooks    string      `json:"webhooks"`
		LastMessage interface{} `json:"last_message"`
	} `json:"links"`
}

// ChannelCreateParams holds information used in creating a new channel.
type ChannelCreateParams struct {
	FriendlyName string          `url:",omitempty"`
	UniqueName   string          `url:",omitempty"`
	Attributes   json.RawMessage `url:",omitempty"`

	// Type can be public or private. Default public.
	Type string `url:",omitempty"`

	// DateCreated ISO-8601 format. Default current time.
	DateCreated string `url:",omitempty"`

	// DateUpdated ISO-8601 format. Default null.
	DateUpdated string `url:",omitempty"`
	// CreatedBy identity of the User that created the channel. Default `system`.
	CreatedBy string `url:",omitempty"`
}

func (c ChannelCreateParams) encode() io.Reader {
	v, _ := query.Values(c)
	return strings.NewReader(v.Encode())
}

// ChannelUpdateParams holds information used in updateing an existing channel.
type ChannelUpdateParams struct {
	FriendlyName string          `url:",omitempty"`
	UniqueName   string          `url:",omitempty"`
	Attributes   json.RawMessage `url:",omitempty"`

	// DateCreated ISO-8601 format.
	DateCreated string `url:",omitempty"`

	// DateUpdated ISO-8601 format.
	DateUpdated string `url:",omitempty"`

	// CreatedBy identity of the User that created the channel. Default `system`.
	CreatedBy string `url:",omitempty"`
}

func (c ChannelUpdateParams) encode() io.Reader {
	v, _ := query.Values(c)
	return strings.NewReader(v.Encode())
}
