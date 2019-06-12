package chat

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/smnalex/twilio-go"
)

// MemberResource handles interactions with Member Programmable Chat REST API.
type MemberResource struct {
	memberAPI
}

// Member represents the membership of a User within the Service instance to a Channel.
type Member struct {
	Sid                      string `json:"sid"`
	AccountSid               string `json:"account_sid"`
	ChannelSid               string `json:"channel_sid"`
	ServiceSid               string `json:"service_sid"`
	Identity                 string `json:"identity"`
	RoleSid                  string `json:"role_sid"`
	LastConsumedMessageIndex int    `json:"last_consumed_message_index"`
	LastConsumptionTimestamp string `json:"last_consumption_timestamp"`

	// DateCreated ISO-8601 format.
	DateCreated string `json:"date_created"`

	// DateUpdated ISO-8601 format.
	DateUpdated string          `json:"date_updated"`
	Attributes  json.RawMessage `json:"attributes"`
	URL         string          `json:"url"`
}

// MemberCreateParams holds information used in adding a member to a channel.
type MemberCreateParams struct {
	Identity                 string
	RoleSid                  string `url:",omitempty"`
	LastConsumedMessageIndex int    `url:",omitempty"`
	LastConsumptionTimestamp string `url:",omitempty"`

	// DateCreated ISO-8601 format.
	DateCreated string `url:",omitempty"`

	// DateUpdated ISO-8601 format.
	DateUpdated string          `url:",omitempty"`
	Attributes  json.RawMessage `url:",omitempty"`
}

func (mcp MemberCreateParams) encode() io.Reader {
	return strings.NewReader(twilio.Values(mcp).Encode())
}
