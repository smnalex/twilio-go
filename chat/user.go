package chat

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/go-querystring/query"
)

// UserResource handles interactions with Users Programmable Chat REST API.
type UserResource struct {
	UserRepository
}

// User resource of Programmable Chat represents a particular user represented by
// an Identity as provided by the developer.
type User struct {
	SID        string `json:"sid"`
	AccountSID string `json:"account_sid"`
	ServiceSID string `json:"service_sid"`
	Identity   string `json:"identity"`
	RoleSID    string `json:"role_sid"`

	// IsOnline is false if the Service's reachability_enabled is false,
	// if the User has never been online for the Service instance, even if the
	// Service's reachability_enabled is true.
	IsOnline bool `json:"is_online"`

	// IsNotifiable is false if the Service's reachability_enabled is false,
	// and if the User has never had a notification registration, even if the
	// Service's reachability_enabled is true.
	IsNotifiable        bool   `json:"is_notifiable"`
	FriendlyName        string `json:"friendly_name"`
	JoinedChannelsCount int    `json:"joined_channels_count"`

	// DateCreated RFC 2822 format.
	DateCreated string `json:"date_created"`

	// DateUpdated RFC 2822 format.
	DateUpdated string `json:"date_updated"`
	Links       struct {
		UserChannels string `json:"user_channels"`
		UserBindings string `json:"user_bindings"`
	} `json:"links"`
	URL        string          `json:"url"`
	Attributes json.RawMessage `json:"attributes,omitempty"`
}

// UserCreateParams holds information used in creating a new user.
// https://www.twilio.com/docs/chat/rest/users#create-a-user
type UserCreateParams struct {
	Identity string
	RoleSID  string `url:"RoleSid,omitempty"`

	// Attributes JSON string that stores application-specific data.
	Attributes   json.RawMessage `url:",omitempty"`
	FriendlyName string          `url:",omitempty"`
}

func (u UserCreateParams) encode() io.Reader {
	b, _ := query.Values(u)
	return strings.NewReader(b.Encode())
}

// UserUpdateParams holds information used in updating an existing user.
// https://www.twilio.com/docs/chat/rest/users#update-a-user
type UserUpdateParams struct {
	RoleSID      string `url:"RoleSid,omitempty"`
	FriendlyName string `url:"FriendlyName,omitempty"`

	// Attributes JSON string that stores application-specific data.
	Attributes json.RawMessage `url:",omitempty"`
}

func (u UserUpdateParams) encode() io.Reader {
	b, _ := query.Values(u)
	return strings.NewReader(b.Encode())
}
