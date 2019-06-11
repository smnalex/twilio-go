package chat

import (
	"io"
	"strings"

	"github.com/google/go-querystring/query"
)

// ServiceResource handles interactions with Services Programmable Chat REST API.
type ServiceResource struct {
	serviceAPI
}

// Service holds a service information.
type Service struct {
	SID                          string                 `json:"sid,omitempty"`
	FriendlyName                 string                 `json:"friendly_name,omitempty"`
	URL                          string                 `json:"url,omitempty"`
	AccountSID                   string                 `json:"account_sid,omitempty"`
	DateCreated                  string                 `json:"date_created,omitempty"`
	DateUpdated                  string                 `json:"date_updated,omitempty"`
	DefaultChannelCreatorRoleSID string                 `json:"default_channel_creator_role_sid,omitempty"`
	DefaultChannelRoleSID        string                 `json:"default_channel_role_sid,omitempty"`
	DefaultServiceRoleSID        string                 `json:"default_service_role_sid,omitempty"`
	PreWebhookURL                string                 `json:"pre_webhook_url,omitempty"`
	PostWebhookURL               string                 `json:"post_webhook_url,omitempty"`
	ConsumptionReportInterval    int                    `json:"consumption_report_interval,omitempty"`
	TypingIndicatorTimeout       int                    `json:"typing_indicator_timeout,omitempty"`
	PreWebhookRetryCount         int                    `json:"pre_webhook_retry_count,omitempty"`
	PostWebhookRetryCount        int                    `json:"post_webhook_retry_count,omitempty"`
	ReachabilityEnabled          bool                   `json:"reachability_enabled,omitempty"`
	ReadStatusEnabled            bool                   `json:"read_status_enabled,omitempty"`
	WebhookFilters               []string               `json:"webhook_filters,omitempty"`
	WebhookMethod                string                 `json:"webhook_method,omitempty"`
	Limits                       map[string]int         `json:"limits,omitempty"`
	Links                        map[string]string      `json:"links,omitempty"`
	Notifications                *Notifications         `json:"notifications,omitempty"`
	Media                        map[string]interface{} `json:"media,omitempty"`
}

// ServiceCreateParams holds information used in creating a new service.
type ServiceCreateParams struct {
	FriendlyName string
}

func (s ServiceCreateParams) encode() io.Reader {
	b, _ := query.Values(s)
	return strings.NewReader(b.Encode())
}

// ServiceUpdateParams holds information used in updating an existing service.
// https://www.twilio.com/docs/chat/rest/services#update-a-service
type ServiceUpdateParams struct {
	FriendlyName                 string         `url:",omitempty"`
	DefaultServiceRoleSID        string         `url:"DefaultServiceRoleSID,omitempty"`
	DefaultChannelCreatorRoleSID string         `url:",omitempty"`
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
	Template string `url:",omitempty"`
	Enabled  bool   `url:",omitempty"`
	Sound    bool   `url:",omitempty"`
}

// Notifications holds a service notification configuration.
type Notifications struct {
	LogEnabled       bool                         `url:",omitempty"`
	AddedToChannel   *NotificationChannelProperty `url:",omitempty"`
	InvitedToChannel *NotificationChannelProperty `url:",omitempty"`
	NewMessage       struct {
		Enabled                      bool `url:",omitempty"`
		*NotificationChannelProperty `url:",omitempty"`
		BadgeCountEnabled            bool `url:",omitempty"`
	} `url:",omitempty"`
	RemoveFromChannel *NotificationChannelProperty `url:",omitempty"`
}

func (s ServiceUpdateParams) encode() io.Reader {
	b, _ := query.Values(s)
	return strings.NewReader(b.Encode())
}
