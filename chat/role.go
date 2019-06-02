package chat

import (
	"io"
	"strings"

	"github.com/google/go-querystring/query"
)

// RoleResource handles interactions with Roles Programmable Chat REST API.
type RoleResource struct {
	RoleRepository
}

// Role holds a role information.
type Role struct {
	SID          string   `json:"sid,omitempty"`
	AccountSID   string   `json:"account_sid,omitempty"`
	ServiceSID   string   `json:"service_sid,omitempty"`
	FriendlyName string   `json:"friendly_name"`
	Type         string   `json:"type"`
	Permissions  []string `json:"permissions"`
	DateCreated  string   `json:"date_created"`
	DateUpdated  string   `json:"date_updated"`
	URL          string   `json:"url"`
}

// RoleCreateParams holds information used in creating a new role.
type RoleCreateParams struct {
	// A descriptive string that you create to describe the new resource.
	// It can be up to 64 characters long.
	FriendlyName string

	// The type of role. Can be: channel for [Channel] (https://www.twilio.com/docs/chat/channels)
	// roles or deployment for [Service](https://www.twilio.com/docs/chat/rest/services) roles.
	Type string

	// https://www.twilio.com/docs/chat/permissions.
	Permission []string
}

func (r RoleCreateParams) encode() io.Reader {
	v, _ := query.Values(r)
	return strings.NewReader(v.Encode())
}

// RoleUpdateParams holds information used in updating an existing role.
type RoleUpdateParams struct {
	// Deployment type: "createChannel", "joinChannel", "editOwnUserInfo",
	// service type: "sendMessage", "leaveChannel", "addMember".
	Permission []string
}

func (r RoleUpdateParams) encode() io.Reader {
	v, _ := query.Values(r)
	return strings.NewReader(v.Encode())
}
