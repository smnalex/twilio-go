package chat

// BindingResource handles interactions with Bindings Programmable Chat REST API.
type BindingResource struct {
	bindingAPI
}

// Binding resource of Programmable Chat represents push notification subscriptions
// for Users within the Service instance.
type Binding struct {
	SID           string   `json:"sid"`
	AccountSid    string   `json:"account_sid"`
	ServiceSid    string   `json:"service_sid"`
	Endpoint      string   `json:"endpoint"`
	Identity      string   `json:"identity"`
	UserSid       string   `json:"user_sid"`
	BindingType   string   `json:"binding_type"`
	CredentialSid string   `json:"credential_sid"`
	MessageTypes  []string `json:"message_types"`
	URL           string   `json:"url"`

	// DateCreated ISO-8601 format.
	DateCreated string `json:"date_created"`

	// DateUpdated ISO-8601 format.
	DateUpdated string `json:"date_updated"`
}
