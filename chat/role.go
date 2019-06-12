package chat

import (
	"io"
	"strings"

	"github.com/smnalex/twilio-go"
)

// RoleResource handles interactions with Roles Programmable Chat REST API.
type RoleResource struct {
	roleAPI
}

// Role represents what a user can do within a Chat Service instance.
// Roles are either Service scoped or Channel scoped.
type Role struct {
	Sid          string   `json:"sid"`
	AccountSid   string   `json:"account_sid"`
	ServiceSid   string   `json:"service_sid"`
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

func (rcp RoleCreateParams) encode() io.Reader {
	return strings.NewReader(twilio.Values(rcp).Encode())
}

// RoleUpdateParams holds information used in updating an existing role.
type RoleUpdateParams struct {
	// Deployment type: "createChannel", "joinChannel", "editOwnUserInfo",
	// service type: "sendMessage", "leaveChannel", "addMember".
	Permission []string
}

func (rup RoleUpdateParams) encode() io.Reader {
	return strings.NewReader(twilio.Values(rup).Encode())
}
