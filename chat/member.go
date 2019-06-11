package chat

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/go-querystring/query"
)

// MemberResource handles interactions with Member Programmable Chat REST API.
type MemberResource struct {
	memberAPI
}

// Member represents the membership of a User within the Service instance to a Channel.
type Member struct {
	Sid                      string          `json:"sid"`
	AccountSID               string          `json:"account_sid"`
	ChannelSID               string          `json:"channel_sid"`
	ServiceSID               string          `json:"service_sid"`
	Identity                 string          `json:"identity"`
	RoleSid                  string          `json:"role_sid"`
	LastConsumedMessageIndex int             `json:"last_consumed_message_index"`
	LastConsumptionTimestamp string          `json:"last_consumption_timestamp"`
	DateCreated              string          `json:"date_created"`
	DateUpdated              string          `json:"date_updated"`
	Attributes               json.RawMessage `json:"attributes"`
	URL                      string          `json:"url"`
}

// MemberCreateParams holds information used in adding a member to a channel.
type MemberCreateParams struct {
	Identity                 string
	RoleSID                  string          `url:"RoleSid,omitempty"`
	LastConsumedMessageIndex int             `url:",omitempty"`
	LastConsumptionTimestamp string          `url:",omitempty"`
	DateCreated              string          `url:",omitempty"`
	DateUpdated              string          `url:",omitempty"`
	Attributes               json.RawMessage `url:",omitempty"`
}

func (mcp MemberCreateParams) encode() io.Reader {
	v, _ := query.Values(mcp)
	return strings.NewReader(v.Encode())
}
