package chat

import (
	"io"
	"strings"
	"time"

	"github.com/smnalex/twilio-go"
)

// InviteResource handles interactions with Invite Programmable Chat REST API.
type InviteResource struct {
	inviteAPI
}

// Invite represents all pending invitations to make Users into Channel Members.
type Invite struct {
	AccountSid  string    `json:"account_sid"`
	ChannelSid  string    `json:"channel_sid"`
	CreatedBy   string    `json:"created_by"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Identity    string    `json:"identity"`
	RoleSid     string    `json:"role_sid"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
}

// InviteCreateParams holds information used in creating a new invite.
type InviteCreateParams struct {
	Identity string
	RoleSid  string `url:",omitempty"`
}

func (icp InviteCreateParams) encode() io.Reader {
	return strings.NewReader(twilio.Values(icp).Encode())
}
