package chat

// ServiceResource handles interactions with Programmable Chat REST API.
type ServiceResource struct {
	ServiceRepository
}

// Service holds a service configuration.
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

// NotificationChannelProperty holds the propeties of push notifications.
type NotificationChannelProperty struct {
	Template string `json:"template,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
	Sound    bool   `json:"sound,omitempty"`
}

// Notifications holds a service notification configuration.
type Notifications struct {
	LogEnabled       bool                         `json:"log_enabled,omitempty"`
	AddedToChannel   *NotificationChannelProperty `json:"added_to_channel,omitempty"`
	InvitedToChannel *NotificationChannelProperty `json:"invited_to_channel,omitempty"`
	NewMessage       struct {
		NotificationChannelProperty
		BadgeCountEnabled bool `json:"badge_count_enabled,omitempty"`
	} `json:"new_message,omitempty"`
	RemoveFromChannel *NotificationChannelProperty `json:"remove_from_channel,omitempty"`
}

// ServiceUpdateParams holds the service update properties.
type ServiceUpdateParams struct {
	FriendlyName                 string                 `json:"friendly_name,omitempty"`
	DefaultServiceRoleSID        string                 `json:"default_service_role_sid,omitempty"`
	DefaultChannelCreatorRoleSID string                 `json:"default_channel_creator_role_sid,omitempty"`
	DefaultChannelRoleSID        string                 `json:"default_channel_role_sid,omitempty"`
	ReadStatusEnabled            bool                   `json:"read_status_enabled,omitempty"`
	ReachabilityEnabled          bool                   `json:"reachability_enabled,omitempty"`
	TypingIndicatorTimeout       int                    `json:"typing_indicator_timeout,omitempty"`
	ConsumptionReportInterval    int                    `json:"consumption_report_interval,omitempty"`
	Notifications                *Notifications         `json:"notifications,omitempty"`
	PreWebhookURL                string                 `json:"pre_webhook_url,omitempty"`
	PostWebhookURL               string                 `json:"post_webhook_url,omitempty"`
	WebhookMethod                string                 `json:"webhook_method,omitempty"`
	WebhookFilters               []string               `json:"webhook_filters,omitempty"`
	Limits                       map[string]int         `json:"limits,omitempty"`
	Media                        map[string]interface{} `json:"media,omitempty"`
	PreWebhookRetryCount         int                    `json:"pre_webhook_retry_count,omitempty"`
	PostWebhookRetryCount        int                    `json:"post_webhook_retry_count,omitempty"`
}
