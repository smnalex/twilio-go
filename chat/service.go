package chat

import (
	"io"
	"strings"

	"github.com/smnalex/twilio-go"
)

// ServiceResource handles interactions with Services Programmable Chat REST API.
type ServiceResource struct {
	serviceAPI
}

// Service holds a service information.
type Service struct {
	Sid                          string                 `json:"sid"`
	FriendlyName                 string                 `json:"friendly_name"`
	URL                          string                 `json:"url"`
	AccountSid                   string                 `json:"account_sid"`
	DateCreated                  string                 `json:"date_created"`
	DateUpdated                  string                 `json:"date_updated"`
	DefaultChannelCreatorRoleSid string                 `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                 `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                 `json:"default_service_role_sid"`
	PreWebhookURL                string                 `json:"pre_webhook_url"`
	PostWebhookURL               string                 `json:"post_webhook_url"`
	ConsumptionReportInterval    int                    `json:"consumption_report_interval"`
	TypingIndicatorTimeout       int                    `json:"typing_indicator_timeout"`
	PreWebhookRetryCount         int                    `json:"pre_webhook_retry_count"`
	PostWebhookRetryCount        int                    `json:"post_webhook_retry_count"`
	ReachabilityEnabled          bool                   `json:"reachability_enabled"`
	ReadStatusEnabled            bool                   `json:"read_status_enabled"`
	WebhookFilters               []string               `json:"webhook_filters"`
	WebhookMethod                string                 `json:"webhook_method"`
	Limits                       map[string]int         `json:"limits"`
	Links                        map[string]string      `json:"links"`
	Notifications                *Notifications         `json:"notifications"`
	Media                        map[string]interface{} `json:"media"`
}

// ServiceCreateParams holds information used in creating a new service.
type ServiceCreateParams struct {
	FriendlyName string
}

func (scp ServiceCreateParams) encode() io.Reader {
	return strings.NewReader(twilio.Values(scp).Encode())
}

// ServiceUpdateParams holds information used in updating an existing service.
// https://www.twilio.com/docs/chat/rest/services#update-a-service
type ServiceUpdateParams struct {
	FriendlyName                 string         `url:",omitempty"`
	DefaultServiceRoleSid        string         `url:",omitempty"`
	DefaultChannelCreatorRoleSid string         `url:",omitempty"`
	ReadStatusEnabled            bool           `url:",omitempty"`
	ReachabilityEnabled          bool           `url:",omitempty"`
	TypingIndicatorTimeout       int            `url:",omitempty"`
	ConsumptionReportInterval    int            `url:",omitempty"`
	Notifications                *Notifications `url:",omitempty"`
	PreWebhookURL                string         `url:",omitempty"`
	PostWebhookURL               string         `url:",omitempty"`
	WebhookMethod                string         `url:",omitempty"`
	WebhookFilters               []string       `url:",omitempty"`
	Limits                       struct {
		ChannelMembers int `url:",omitempty"`
		UserChannels   int `url:",omitempty"`
	} `url:",omitempty"`
	Media struct {
		CompatibilityMessage string `url:",omitempty"`
	} `url:",omitempty"`
	PreWebhookRetryCount  int `url:",omitempty"`
	PostWebhookRetryCount int `url:",omitempty"`
}

// NotificationChannelProperty holds the propeties of push notifications.
type NotificationChannelProperty struct {
	Template          string `url:",omitempty"`
	Enabled           bool   `url:",omitempty"`
	Sound             bool   `url:",omitempty"`
	BadgeCountEnabled bool   `url:",omitempty"`
}

// Notifications holds a service notification configuration.
type Notifications struct {
	LogEnabled        bool                         `url:",omitempty"`
	AddedToChannel    *NotificationChannelProperty `url:",omitempty"`
	InvitedToChannel  *NotificationChannelProperty `url:",omitempty"`
	NewMessage        *NotificationChannelProperty `url:",omitempty"`
	RemoveFromChannel *NotificationChannelProperty `url:",omitempty"`
}

func (sup ServiceUpdateParams) encode() io.Reader {
	return strings.NewReader(twilio.Values(sup).Encode())
}
