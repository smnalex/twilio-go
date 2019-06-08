package chat

import (
	"time"
)

// BindingResource handles interactions with Bindings Programmable Chat REST API.
type BindingResource struct {
	BindingRepository
}

// Binding ...
type Binding struct {
	SID           string    `json:"sid"`
	AccountSid    string    `json:"account_sid"`
	ServiceSid    string    `json:"service_sid"`
	DateCreated   time.Time `json:"date_created"`
	DateUpdated   time.Time `json:"date_updated"`
	Endpoint      string    `json:"endpoint"`
	Identity      string    `json:"identity"`
	UserSid       string    `json:"user_sid"`
	BindingType   string    `json:"binding_type"`
	CredentialSid string    `json:"credential_sid"`
	MessageTypes  []string  `json:"message_types"`
	URL           string    `json:"url"`
}
