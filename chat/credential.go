package chat

import (
	"io"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// CredentialResource handles interactions with Credential Programmable Chat REST API.
type CredentialResource struct {
	credentialAPI
}

// Credential resource represents one credential record for a particular push notifications channel.
// Currently APNS, FCM and GCM types are supported.
type Credential struct {
	Sid          string    `json:"sid"`
	AccountSid   string    `json:"account_sid"`
	FriendlyName string    `json:"friendly_name"`
	Type         string    `json:"type"`
	Sandbox      string    `json:"sandbox"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	URL          string    `json:"url"`
}

// CredentialCreateParams holds information used in creating a new credential.
// https://www.twilio.com/docs/chat/rest/credentials#create-a-credential
type CredentialCreateParams struct {
	Type         string
	FriendlyName string `url:",omitempty"`
	Certificate  string `url:",omitempty"`
	PrivateKey   string `url:",omitempty"`
	Sandbox      bool   `url:",omitempty"`
	APIKey       string `url:"ApiKey,omitempty"`
	Secret       string `url:",omitempty"`
}

func (ccp CredentialCreateParams) encode() io.Reader {
	v, _ := query.Values(ccp)
	return strings.NewReader(v.Encode())
}

// CredentialUpdateParams holds information used in updateing an existing credential.
// https://www.twilio.com/docs/chat/rest/credentials#update-a-credential
type CredentialUpdateParams struct {
	FriendlyName string `url:",omitempty"`
	Certificate  string `url:",omitempty"`
	PrivateKey   string `url:",omitempty"`
	Sandbox      bool   `url:",omitempty"`
	APIKey       string `url:"ApiKey,omitempty"`
	Secret       string `url:",omitempty"`
}

func (ccp CredentialUpdateParams) encode() io.Reader {
	v, _ := query.Values(ccp)
	return strings.NewReader(v.Encode())
}
