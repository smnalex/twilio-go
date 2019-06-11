package chat

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/go-querystring/query"
)

// MessageResource handles interactions with Messages Programmable Chat REST API.
type MessageResource struct {
	messageAPI
}

// Message represents a single message within a Channel
// in a Service instance.
type Message struct {
	SID        string `json:"sid"`
	AccountSID string `json:"account_sid"`
	ServiceSID string `json:"service_sid"`
	ChannelSID string `json:"channel_sid"`
	From       string `json:"from"`
	To         string `json:"to"`

	// DateCreated ISO8601 format
	DateCreated string `json:"date_created"`

	// DateUpdated ISO8601 format
	DateUpdated   string          `json:"date_updated"`
	LastUpdatedBy string          `json:"last_updated_by"`
	WasEdited     bool            `json:"was_edited"`
	Body          string          `json:"body"`
	Index         int             `json:"index"`
	Type          string          `json:"type"`
	Media         interface{}     `json:"media"`
	URL           string          `json:"url"`
	Attributes    json.RawMessage `json:"attributes"`
}

// MessageCreateParams holds information used in sending a new message.
type MessageCreateParams struct {
	From string `url:",omitempty"`

	// DateCreated ISO8601 format
	DateCreated string `url:",omitempty"`

	// DateUpdated ISO8601 format
	DateUpdated   string          `url:",omitempty"`
	LastUpdatedBy string          `url:",omitempty"`
	Body          string          `url:",omitempty"`
	MediaSid      string          `url:",omitempty"`
	Attributes    json.RawMessage `url:",omitempty"`
}

func (mcp MessageCreateParams) encode() io.Reader {
	v, _ := query.Values(mcp)
	return strings.NewReader(v.Encode())
}

// MessageUpdateParams holds information used in updating an existing message.
type MessageUpdateParams struct {
	From string `url:",omitempty"`
	Body string `url:",omitempty"`
	// DateCreated ISO8601 format
	DateCreated string `url:",omitempty"`

	// DateUpdated ISO8601 format
	DateUpdated   string          `url:",omitempty"`
	LastUpdatedBy string          `url:",omitempty"`
	Attributes    json.RawMessage `url:",omitempty"`
}

func (mup MessageUpdateParams) encode() io.Reader {
	v, _ := query.Values(mup)
	return strings.NewReader(v.Encode())
}
